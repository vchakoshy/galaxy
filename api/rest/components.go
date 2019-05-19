package rest

import (
	"context"
	"encoding/json"
	"log"

	"github.com/davecgh/go-spew/spew"
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
		cs := flexComponentSettings{}
		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)

		if cs.Settings.DataProvider == "BOOK" {
			queries := cs.Settings.Setup.getQueries()
			queries = append(queries, qm.Limit(8))

			spew.Dump(cs.Elements.Title)
			com := flexComponent{
				Type:         "HL_BOOKS_ARTICLE",
				ResourceType: "BOOK",
				Title:        cs.Elements.Title.Value.Static,
			}

			queries = append(queries, cs.Settings.Setup.getSort()...)

			log.Println(queries)

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
