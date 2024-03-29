package flex

import (
	"context"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

const (
	dataProviderTypePublisher = "PUBLISHER"
)

// PublisherDataProvider struct
type PublisherDataProvider struct {
	ComponentSettings ComponentSettings
	ComponentType     string
}

func NewPublisherDataProvider(cs ComponentSettings) PublisherDataProvider {
	return PublisherDataProvider{
		ComponentSettings: cs,
		ComponentType:     "publisher",
	}
}

func (b PublisherDataProvider) GetOutputComponent() OutputComponent {
	return b.getOutputComponent()
}

func (b PublisherDataProvider) getOutputComponent() OutputComponent {
	com, _ := NewOutputComponent(
		OutputComponentSetTitle(b.ComponentSettings.Elements.Title.Value.Static),
		OutputComponentSetType(b.ComponentType),
		OutputComponentSetResourceType(dataProviderTypePublisher),
	)

	// com := OutputComponent{
	// 	Type:         b.ComponentType,
	// 	ResourceType: dataProviderTypePublisher,
	// 	Title:        b.ComponentSettings.Elements.Title.Value.Static,
	// }

	a := b.ComponentSettings.Elements.MoreTitle.Action.getAction()
	if a.Type != "" {
		com.Action = a
	}

	if b.ComponentSettings.Elements.MoreTitle.Value != "" {
		com.ActionTitle = b.ComponentSettings.Elements.MoreTitle.Value
	}

	p, err := b.Models()
	if err != nil {
		log.Println(err.Error())
		return *com
	}

	publishers := p.(models.PublisherSlice)

	for i, v := range publishers {
		com.Data.Items.Generic[i] = Generic{
			Title: v.Title,
			Image: v.Cover.String,
			Action: &Action{ // TODO fix actions
				Type:   "publisher_page",
				Method: "/general/profile/publisher/get",
				Input: []InputAction{
					InputAction{
						Key:   "publisherId",
						Value: v.ID,
					},
					InputAction{
						Key:   "pageName",
						Value: "PUBLISHER_PAGE",
					},
				},
			},
		}

		// com.Data.Items.Model = append(com.Data.Items.Model, v)
	}

	return *com
}

func (b PublisherDataProvider) Models() (r interface{}, err error) {
	queries := []qm.QueryMod{}

	queries = append(queries, qm.Where("content_provider_type=?", "BOOK"))

	publisherIDs := b.ComponentSettings.Settings.Setup.Publisher.GetIdis()
	if len(publisherIDs) > 0 {
		queries = append(queries, qm.WhereIn("publisher.id in ?", publisherIDs...))
	}

	channelIDs := b.ComponentSettings.Settings.Setup.ProposedList.GetIdis()
	if len(channelIDs) > 0 {
		queries = append(queries, qm.InnerJoin("channel ON channel.publisher_id = publisher.id"))
		queries = append(queries, qm.WhereIn("channel.id in ?", channelIDs...))
	}

	queries = append(queries, qm.Limit(10))

	r, err = models.Publishers(queries...).AllG(context.Background())

	return
}
