package flex

import (
	"context"
	"log"
	"strings"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/hubble"
)

const (
	sortTypeRecent         = "RECENT"
	sortTypeRecentPublish  = "RECENT_PUBLISH"
	sortTypeLowestPrice    = "LOWEST_PRICE"
	sortTypeBestSellect    = "BESTSELLER"
	sortTypeWeekBestSeller = "WEEK_BESTSELLER"
	sortTypePopular        = "POPULAR"
	sortTypeMostCommented  = "MOST_COMMENTED"
)

type DataProvidersBook struct{}

func (b DataProvidersBook) getGeneric(cs ComponentSettings, t string) OutputComponent {
	com := OutputComponent{
		Type:         t,
		ResourceType: "BOOK",
		Title:        cs.Elements.Title.Value.Static,
	}

	a := getAction(cs.Elements.MoreTitle.Action)
	if a.Type != "" {
		com.Action = a
	}

	if cs.Elements.MoreTitle.Value != "" {
		com.ActionTitle = cs.Elements.MoreTitle.Value
	}

	ss := esClient.Search(hubble.ProductIndexName)

	switch cs.Settings.Setup.Sort.Value {
	case sortTypeRecent:
		ss.Sort("id", false)
	case sortTypeRecentPublish:
		ss.Sort("publish_date", false)
	case sortTypeLowestPrice:
		ss.Sort("price", true)
	case sortTypeBestSellect:
		ss.Sort("all_sales_count", false)
	case sortTypeWeekBestSeller:
		ss.Sort("week_sales_count", false)
	case sortTypePopular:
		ss.Sort("month_download_count", false)
	case sortTypeMostCommented:
		ss.Sort("all_comment_count", false)
	default:
		ss.Sort("publish_time", false)
	}

	q := cs.Settings.Setup.GetQueries()

	esres, err := ss.
		StoredFields("id").
		Query(q).
		Size(8).
		Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return OutputComponent{}
	}

	bookIdis := QueryIdis{}

	for _, item := range esres.Hits.Hits {
		itemID := strings.Replace(item.Id, "book-", "", 1)
		bookIdis = append(bookIdis, itemID)
	}

	queries := []qm.QueryMod{}

	if len(bookIdis) > 0 {
		queries = append(queries, qm.WhereIn("id in ?", bookIdis...))
	}

	queries = append(queries, qm.Limit(8))

	res, modelRes := newGenericBookByQuery(queries)

	com.Data.Items.Generic = make([]Generic, len(res))
	com.Data.Items.Model = make([]Book, len(res))
	for i, v := range res {
		com.Data.Items.Generic[i] = v
		com.Data.Items.Model[i] = modelRes[i]
	}

	return com
}
