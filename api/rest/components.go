package rest

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

// Component for flex
type Component struct {
	queries []qm.QueryMod
	pageID  int
}

func NewComponentByPage(pageID int) *Component {
	return &Component{
		pageID: pageID,
	}
}

func (c *Component) getData() []flexComponent {
	qComponent := qm.Where("page_id=?", c.pageID)

	fpc, err := models.FlexPageComponents(
		qComponent,
		qm.OrderBy("crud_order asc")).
		AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	componenets := make([]flexComponent, 0)

	for _, comp := range fpc {

		compModel, err := models.FlexComponents(qm.Where("id=?", comp.ComponentID)).OneG(context.Background())
		if err != nil {
			log.Println(err.Error(), comp.ComponentID)
		}

		cs := flexComponentSettings{}
		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)

		if cs.Settings.DataProvider == "BOOK" {
			queries := cs.Settings.Setup.getQueries()
			queries = append(queries, qm.Limit(8))

			com := flexComponent{
				Type:         compModel.Type,
				ResourceType: "BOOK",
				Title:        cs.Elements.Title.Value.Static,
				Action: flexBaseAction{
					Type: strings.ToLower(cs.Elements.MoreTitle.Action.Type),
					Input: []flexComponentAction{
						{
							Key:        "categoryId",
							ArrayValue: []string{"10036", "10221"},
						},
						{
							Key:   "free",
							Value: false,
						},
						{
							Key:   "subscription",
							Value: false,
						},
					},
					ExtraData: nil,
					Method:    "/v2/general/list/book",
				},
			}

			if cs.Elements.MoreTitle.Value != "" {
				com.ActionTitle = cs.Elements.MoreTitle.Value
			}

			queries = append(queries, cs.Settings.Setup.getSort()...)

			// log.Println(queries)

			res := newGenericBookByQuery(queries)

			com.Data.Items.Generic = make([]interface{}, len(res))
			for i, v := range res {
				com.Data.Items.Generic[i] = v
			}

			componenets = append(componenets, com)
		}
	}
	return componenets
}
