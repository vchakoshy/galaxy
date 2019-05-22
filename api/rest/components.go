package rest

import (
	"context"
	"encoding/json"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/flex"
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

		compModel, err := models.
			FlexComponents(qm.Where("id=?", comp.ComponentID)).
			OneG(context.Background())
		if err != nil {
			log.Println(err.Error(), comp.ComponentID)
			continue
		}

		cs := flexComponentSettings{}
		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)
		log.Println(cs.Settings.DataProvider)

		if cs.Settings.DataProvider == "BOOK" {
			com := flexComponent{
				Type:         compModel.Type,
				ResourceType: "BOOK",
				Title:        cs.Elements.Title.Value.Static,
				Action:       getAction(cs),
			}

			if cs.Elements.MoreTitle.Value != "" {
				com.ActionTitle = cs.Elements.MoreTitle.Value
			}

			queries := []qm.QueryMod{}

			inList := cs.Settings.Setup.Format.Value.getInterfaceList()
			if len(inList) > 0 {
				queries = append(queries, qm.WhereIn("format in ?", inList...))
			}
			// queries := cs.Settings.Setup.getQueries()
			queries = append(queries, qm.Limit(8))
			queries = append(queries, cs.Settings.Setup.getSort()...)

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

func getAction(cs flexComponentSettings) flexBaseAction {
	action := flex.GetActionByPanelType(cs.Elements.MoreTitle.Action.Type)

	return flexBaseAction{
		Type:      action.ClientType,
		Input:     cs.Settings.Setup.getInputActions(),
		ExtraData: nil,
		Method:    action.Method,
	}
}
