package rest

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
			set := cs.Settings.Setup.Book

			queries := make([]qm.QueryMod, 0)

			qidis := make([]interface{}, len(set.Ids))
			for _, id := range set.Ids {
				qidis = append(qidis, id)
			}

			if len(qidis) > 0 {
				queries = append(queries, qm.WhereIn("id in ?", qidis...))
			}

			queries = append(queries, qm.Limit(8))

			com := flexComponent{
				Type:         "HL_BOOKS_ARTICLE",
				ResourceType: "BOOK",
			}

			log.Println(cs.Settings.Setup.Sort)

			switch cs.Settings.Setup.Sort.Value {
			case "RECENT":
				queries = append(queries, qm.OrderBy("id DESC"))
			}

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
