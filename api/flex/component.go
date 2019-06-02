package flex

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

func (c *Component) GetData(pageid int) []OutputComponent {
	itemsPerPage := 5
	offset := (pageid - 1) * itemsPerPage

	fpc, err := models.FlexPageComponents(
		qm.Where("page_id=?", c.pageID),
		qm.And("deleted_at IS NULL"),
		qm.And("active = ?", 1),
		qm.Limit(itemsPerPage),
		qm.Offset(offset),
		qm.OrderBy("crud_order asc")).
		AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	componenets := make([]OutputComponent, 0)

	for _, comp := range fpc {

		compModel, err := models.
			FlexComponents(qm.Where("id=?", comp.ComponentID)).
			OneG(context.Background())
		if err != nil {
			log.Println(err.Error(), comp.ComponentID)
			continue
		}

		cs := ComponentSettings{}

		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)

		switch compModel.Type {
		case "HL_BOOKS_FULL", "HL_BOOKS_ARTICLE", "HL_BOOKS_WIDE",
			"HL_BOOKS_MAGAZINE", "HL_CARDS_FULL", "HL_CARDS_MINI",
			"HL_LOGOS", "VL_CARDS_BIG", "VL_CARDS_FULL", "VL_CONTENT_LIST_MORE",
			"HL_TAGS", "VL_CARDS_ARROWED", "VL_CARDS_BOOKMARKABLE",
			"VL_CARDS_SIMPLE_MORE", "VL_CARDS_FULL_MORE", "HL_GRID_SIMPLE",
			"VL_BOOKS_DESCRIPTION":
			componenets = append(componenets, handleListComponent(cs, compModel.Type))
		case "SINGLE_BANNER_SIMPLE", "SINGLE_BUTTON_SIMPLE", "SINGLE_TEXT_LINK",
			"SINGLE_TEXT_SIMPLE", "SERVER_BOOK_REQUEST":
			componenets = append(componenets, handleSingleComponent(cs, compModel.Type))
		case "HL_BOOKS_LIBRARY":
			componenets = append(componenets, handleBooksLibraryComponent(cs))
		default:
			log.Println("Unhandled: " + compModel.Type)
		}
	}
	return componenets
}

const (
	actionContentListPanelType      = "CONTENT_LIST"
	actionBookPanelType             = "BOOK_PAGE"
	actionProposedListPagePanelType = "PROPOSED_LIST_PAGE"
	actionPublisherListPanelType    = "PUBLISHER_LIST"
	actionNewsListPanelType         = "NEWS_LIST"
)

func (acs ActionCS) getAction() (a *Action) {
	a = &Action{
		Type:      strings.ToLower(acs.Type),
		Input:     acs.Setup.getInputActions(),
		ExtraData: nil,
	}

	switch acs.Type {
	case actionContentListPanelType:
		a.Method = "/v2/general/list/book"
	case actionBookPanelType:
		a.Method = "/"
	case actionProposedListPagePanelType:
		a.Method = "/general/proposed-list/get"
	case actionPublisherListPanelType:
		a.Method = "/v2/general/list/publisher"
	case actionNewsListPanelType:
		a.Method = "/v2/general/list/news"
	default:
		a.Method = "/"
	}

	return
}
