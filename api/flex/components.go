package flex

import (
	"context"
	"log"
	"strings"

	"gitlab.fidibo.com/backend/galaxy/hubble"

	// "github.com/davecgh/go-spew/spew"
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

		bookIdis := QueryIdis{}

		ss := esClient.Search(hubble.ProductIndexName)

		switch cs.Settings.Setup.Sort.Value {
		case "BESTSELLER":
			ss.Sort("all_sales_count", false)

		case "WEEK_BESTSELLER":
			ss.Sort("week_sales_count", false)

		case "POPULAR":
			ss.Sort("month_download_count", false)

		case "RECENT":
			ss.Sort("publish_time", false)

		case "MOST_COMMENTED":
			ss.Sort("all_comment_count", false)

		default:
			ss.Sort("publish_time", false)
		}

		q := elastic.NewBoolQuery()

		filters := elastic.NewBoolQuery()

		formatList := cs.Settings.Setup.Format.Value.getInterfaceList()
		// spew.Dump(formatList)
		if len(formatList) > 0 {
			filters.Should(
				elastic.NewTermsQuery("format", formatList...),
			)
		}

		contentTypeList := cs.Settings.Setup.ContentType.Value.getInterfaceList()
		if len(contentTypeList) > 0 {
			filters.Should(
				elastic.NewTermsQuery("content_type", contentTypeList...),
			)
		}

		q.Must(filters)

		esres, err := ss.
			StoredFields("_id").
			Query(q).
			Size(8).
			Do(context.Background())
		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, item := range esres.Hits.Hits {
			itemID := strings.Replace(item.Id, "book-", "", 1)
			bookIdis = append(bookIdis, itemID)
		}

		queries := []qm.QueryMod{}

		if len(bookIdis) > 0 {
			queries = append(queries, qm.WhereIn("id in ?", bookIdis...))
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

func handleBooksLibraryComponent(cs flexComponentSettings) (com flexComponent) {
	m := make(map[string]string)
	m["format"] = cs.Settings.Format
	m["content_type"] = cs.Settings.ContentType
	com.Data.Items.Generic = append(com.Data.Items.Generic, m)
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = "HL_BOOKS_LIBRARY"
	com.ResourceType = "CUSTOM"
	return
}
