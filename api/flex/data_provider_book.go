package flex

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
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

const (
	dataProviderTypeBook = "BOOK"

	elasticFieldID                 = "id"
	elasticFieldPublishDate        = "publish_date"
	elasticFieldPrice              = "price"
	elasticFieldAllSalesCount      = "all_sales_count"
	elasticFieldWeekSalesCount     = "week_sales_count"
	elasticFieldMonthDownloadCount = "month_download_count"
	elasticFieldAllCommentCount    = "all_comment_count"
	elasticFieldPublishTime        = "publish_time"
)

// BookDataProvider struct
type BookDataProvider struct {
	ComponentSettings ComponentSettings
	ComponentType     string
}

func (b BookDataProvider) GetOutputComponent() OutputComponent {
	return b.getOutputComponent()
}

func (b BookDataProvider) getOutputComponent() OutputComponent {
	com := OutputComponent{
		Type:         b.ComponentType,
		ResourceType: dataProviderTypeBook,
		Title:        b.ComponentSettings.Elements.Title.Value.Static,
	}

	a := b.ComponentSettings.Elements.MoreTitle.Action.getAction()
	if a.Type != "" {
		com.Action = a
	}

	if b.ComponentSettings.Elements.MoreTitle.Value != "" {
		com.ActionTitle = b.ComponentSettings.Elements.MoreTitle.Value
	}

	bs, err := b.Models()
	if err != nil {
		log.Println(err.Error())
		return com
	}

	com.Data.Items.Generic = make([]Generic, len(bs))
	com.Data.Items.Model = make([]interface{}, len(bs))
	for i, bk := range bs {
		com.Data.Items.Generic[i] = newGenericByModel(bk)
		com.Data.Items.Model[i] = newBookByModel(bk)

	}

	return com
}

func (b BookDataProvider) Models() (r models.BookSlice, err error) {
	ss := esClient.Search(hubble.ProductIndexName)

	switch b.ComponentSettings.Settings.Setup.Sort.Value {
	case sortTypeRecent:
		ss.Sort(elasticFieldID, false)
	case sortTypeRecentPublish:
		ss.Sort(elasticFieldPublishDate, false)
	case sortTypeLowestPrice:
		ss.Sort(elasticFieldPrice, true)
	case sortTypeBestSellect:
		ss.Sort(elasticFieldAllSalesCount, false)
	case sortTypeWeekBestSeller:
		ss.Sort(elasticFieldWeekSalesCount, false)
	case sortTypePopular:
		ss.Sort(elasticFieldMonthDownloadCount, false)
	case sortTypeMostCommented:
		ss.Sort(elasticFieldAllCommentCount, false)
	default:
		ss.Sort(elasticFieldPublishTime, false)
	}

	q := b.ComponentSettings.Settings.Setup.GetQueries()

	// TODO: enable this in future
	// q.Must(elastic.NewTermQuery("publish", 1))

	esres, err := ss.
		StoredFields("id").
		Query(q).
		Size(8).
		Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		return
	}

	bookIdis := QueryIdis{}

	for _, item := range esres.Hits.Hits {
		itemID := strings.Replace(item.Id, "book-", "", 1)
		bookIdis = append(bookIdis, itemID)
	}

	ids := bookIdis.getInts()

	iids := make([]interface{}, len(ids))
	idStrs := make([]string, len(ids))
	for index, num := range ids {
		iids[index] = num
		idStrs[index] = strconv.Itoa(num)
	}

	orderByIdid := fmt.Sprintf("FIELD(id,%s)", strings.Join(idStrs, ","))

	r, err = models.Books(
		qm.WhereIn("id in ?", iids...),
		qm.OrderBy(orderByIdid),
		qm.Load("Publisher"),
		qm.Load("Author")).
		AllG(context.Background())

	return
}
