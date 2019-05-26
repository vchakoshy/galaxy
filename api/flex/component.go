package flex

import (
	"context"
	"encoding/json"
	"log"

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

func (c *Component) GetData() []flexComponent {
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

		switch compModel.Type {
		case "HL_BOOKS_FULL", "HL_BOOKS_ARTICLE", "HL_BOOKS_WIDE",
			"HL_BOOKS_MAGAZINE", "HL_CARDS_FULL", "HL_CARDS_MINI",
			"HL_LOGOS", "VL_CARDS_BIG", "VL_CARDS_FULL", "VL_CONTENT_LIST_MORE",
			"HL_TAGS", "VL_CARDS_ARROWED", "VL_CARDS_BOOKMARKABLE",
			"VL_CARDS_SIMPLE_MORE", "VL_CARDS_FULL_MORE", "HL_GRID_SIMPLE",
			"VL_BOOKS_DESCRIPTION":
			componenets = append(componenets, handleListComponent(cs, compModel.Type))
		}
	}
	return componenets
}

func getAction(cs flexComponentSettings) flexBaseAction {
	action := GetActionByPanelType(cs.Elements.MoreTitle.Action.Type)

	return flexBaseAction{
		Type:      action.ClientType,
		Input:     cs.Elements.MoreTitle.Action.Setup.getInputActions(),
		ExtraData: nil,
		Method:    action.Method,
	}
}
