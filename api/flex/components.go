package flex

import (
	"log"

	"github.com/olivere/elastic"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func handleListComponent(cs flexComponentSettings, t string) (com flexComponent) {
	if cs.Settings.DataProvider == "BOOK" {
		com = flexComponent{
			Type:         t,
			ResourceType: "BOOK",
			Title:        cs.Elements.Title.Value.Static,
			Action:       getAction(cs),
		}

		if cs.Elements.MoreTitle.Value != "" {
			com.ActionTitle = cs.Elements.MoreTitle.Value
		}

		queries := []qm.QueryMod{}

		bookIdis := QueryIdis{}

		ss := elastic.NewSearchService(esClient)

		switch cs.Settings.Setup.Sort.Value {
		default:
			log.Println("not implemented:", cs.Settings.Setup.Sort.Value)
		case "BESTSELLER":
			ss.Sort("all_sales_count", false)

		case "WEEK_BESTSELLER":
			ss.Sort("week_sales_count", false)

		case "POPULAR":
			ss.Sort("month_download_count", false)

		case "MOST_COMMENTED":
			// Not implemented, this functionallity has performance concern

			// 	q = append(q, qm.InnerJoin("comment ON comment.book_id = book.id"))
			// 	q = append(q, qm.Where("comment.publish = 1"))
			// 	q = append(q, qm.OrderBy("COUNT(comment.id) DESC"))
		}

		// for _, bs := range d {
		// 	bookIdis = append(bookIdis, bs.BookID)
		// }

		log.Println("idis:", bookIdis)

		if len(bookIdis) > 0 {
			queries = append(queries, qm.WhereIn("id in ?", bookIdis...))
		}

		inList := cs.Settings.Setup.Format.Value.getInterfaceList()
		if len(inList) > 0 {
			elastic.NewBoolQuery().Should(
				elastic.NewTermsQuery("format", inList),
			)
			// queries = append(queries, qm.WhereIn("format in ?", inList...))
		}

		queries = append(queries, qm.Limit(8))

		res := newGenericBookByQuery(queries)

		com.Data.Items.Generic = make([]interface{}, len(res))
		for i, v := range res {
			com.Data.Items.Generic[i] = v
		}

	}
	return
}
