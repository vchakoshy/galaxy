package flex

import (
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

type PublisherDataProvider struct{}

func (b PublisherDataProvider) getOutputComponent(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: "PUBLISHER",
		Title:        cs.Elements.Title.Value.Static,
	}

	a := getAction(cs.Elements.MoreTitle.Action)
	if a.Type != "" {
		com.Action = a
	}

	if cs.Elements.MoreTitle.Value != "" {
		com.ActionTitle = cs.Elements.MoreTitle.Value
	}

	queries := []qm.QueryMod{}

	qm.Where("content_provider_type='BOOK'")

	publisherIDs := cs.Settings.Setup.Publisher.GetIdis()
	if len(publisherIDs) > 0 {
		qm.WhereIn("publisher.id", publisherIDs)
	}

	channelIDs := cs.Settings.Setup.ProposedList.GetIdis()
	if len(channelIDs) > 0 {
		qm.InnerJoin("channel ON channel.publisher_id = publisher.id")
		qm.WhereIn("channel.id", channelIDs)
	}

	queries = append(queries, qm.Limit(10))

	publishers, err := models.Publishers(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
		return com
	}

	com.Data.Items.Generic = make([]Generic, len(publishers))
	for i, v := range publishers {
		com.Data.Items.Generic[i] = Generic{
			Title: v.Title,
			Image: v.Cover.String,
			Action: &BaseAction{ // TODO fix actions
				Type:   "publisher_page",
				Method: "/general/profile/publisher/get",
				Input: []ComponentAction{
					ComponentAction{
						Key:   "publisherId",
						Value: v.ID,
					},
					ComponentAction{
						Key:   "pageName",
						Value: "PUBLISHER_PAGE",
					},
				},
			},
		}

		com.Data.Items.Model = append(com.Data.Items.Model, v)
	}

	return com
}
