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

			bookIdis := QueryIdis{}

			switch cs.Settings.Setup.Sort.Value {
			case "BESTSELLER":
				qs := []qm.QueryMod{}
				qs = append(qs, qm.OrderBy("all_sales_count DESC"))
				qs = append(qs, qm.Limit(10))
				d, err := models.BookStats(qs...).AllG(context.Background())
				if err != nil {
					log.Println(err.Error())
					break
				}
				for _, bs := range d {
					bookIdis = append(bookIdis, bs.BookID)
				}

			case "WEEK_BESTSELLER":
				qs := []qm.QueryMod{}
				qs = append(qs, qm.OrderBy("week_sales_count DESC"))
				qs = append(qs, qm.Limit(10))
				d, err := models.BookStats(qs...).AllG(context.Background())
				if err != nil {
					log.Println(err.Error())
					break
				}
				for _, bs := range d {
					bookIdis = append(bookIdis, bs.BookID)
				}

			case "POPULAR":
				qs := []qm.QueryMod{}
				qs = append(qs, qm.OrderBy("month_download_count DESC"))
				qs = append(qs, qm.Limit(10))
				d, err := models.BookStats(qs...).AllG(context.Background())
				if err != nil {
					log.Println(err.Error())
					break
				}
				for _, bs := range d {
					bookIdis = append(bookIdis, bs.BookID)
				}

			case "MOST_COMMENTED":
				// Not implemented, this functionallity has performance concern

				// 	q = append(q, qm.InnerJoin("comment ON comment.book_id = book.id"))
				// 	q = append(q, qm.Where("comment.publish = 1"))
				// 	q = append(q, qm.OrderBy("COUNT(comment.id) DESC"))
			}

			log.Println("idis:", bookIdis)

			if len(bookIdis) > 0 {
				queries = append(queries, qm.WhereIn("id in ?", bookIdis...))
			}

			inList := cs.Settings.Setup.Format.Value.getInterfaceList()
			if len(inList) > 0 {
				queries = append(queries, qm.WhereIn("format in ?", inList...))
			}

			queries = append(queries, qm.Limit(8))

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
	action := GetActionByPanelType(cs.Elements.MoreTitle.Action.Type)

	return flexBaseAction{
		Type:      action.ClientType,
		Input:     cs.Elements.MoreTitle.Action.Setup.getInputActions(),
		ExtraData: nil,
		Method:    action.Method,
	}
}
