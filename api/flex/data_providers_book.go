package flex

import (
	"context"
	"log"
	"strings"

	"github.com/olivere/elastic"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/hubble"
)

type DataProvidersBook struct{}

func (b DataProvidersBook) getGeneric(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: "BOOK",
		Title:        cs.Elements.Title.Value.Static,
	}

	a := getAction(cs)
	if a.Type != "" {
		com.Action = a
	}

	if cs.Elements.MoreTitle.Value != "" {
		com.ActionTitle = cs.Elements.MoreTitle.Value
	}

	bookIdis := QueryIdis{}

	ss := esClient.Search(hubble.ProductIndexName)

	switch cs.Settings.Setup.Sort.Value {

	case "RECENT":
		ss.Sort("id", false)

	case "RECENT_PUBLISH":
		ss.Sort("publish_date", false)

	case "LOWEST_PRICE":
		ss.Sort("price", true)

	case "BESTSELLER":
		ss.Sort("all_sales_count", false)

	case "WEEK_BESTSELLER":
		ss.Sort("week_sales_count", false)

	case "POPULAR":
		ss.Sort("month_download_count", false)

	case "MOST_COMMENTED":
		ss.Sort("all_comment_count", false)

	default:
		ss.Sort("publish_time", false)
	}

	q := elastic.NewBoolQuery()

	catIds := cs.Settings.Setup.Category.GetIdis()
	if len(catIds) > 0 {
		q.Must(
			elastic.NewTermsQuery("categories.id", catIds...),
		)
	}

	formatList := cs.Settings.Setup.Format.Value.getInterfaceList()
	if len(formatList) > 0 {
		q.Must(
			elastic.NewTermsQuery("format.keyword", formatList...),
		)
	}

	contentTypeList := cs.Settings.Setup.ContentType.Value.getInterfaceStringLowerList()
	if len(contentTypeList) > 0 {
		q.Must(
			elastic.NewTermsQuery("content_type.keyword", contentTypeList...),
		)
	}

	esres, err := ss.
		StoredFields("id").
		Query(q).
		Size(8).
		Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return OutputComponent{}
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

	return com
}
