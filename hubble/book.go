package hubble

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

// Book struct of elasticsearch
type Book struct {
	ID                 int        `json:"id"`
	Title              string     `json:"title"`
	Free               int8       `json:"free"`
	Publish            int8       `json:"publish"`
	PublishTime        time.Time  `json:"publish_time"`
	Format             string     `json:"format"`
	Language           string     `json:"language"`
	ContentType        string     `json:"content_type"`
	Price              int        `json:"price"`
	AllCommentCount    int        `json:"all_comment_count,omitempty"`
	AllDownloadCount   int        `json:"all_download_count"`
	AllSalesCount      int        `json:"all_sales_count"`
	DayDownloadCount   int        `json:"day_download_count"`
	DaySalesCount      int        `json:"day_sales_count,omitempty"`
	MonthDownloadCount int        `json:"month_download_count"`
	MonthSalesCount    int        `json:"month_sales_count,omitempty"`
	WeekDownloadCount  int        `json:"week_download_count"`
	WeekSalesCount     int        `json:"week_sales_count"`
	Publisher          Publisher  `json:"publisher"`
	Categories         []Category `json:"categories"`
	Authors            []Author   `json:"authors"`
	Translators        []Author   `json:"translator,omitempty"`

	m *models.Book
}

func (e *Book) addAuthor(d *models.Author) {
	e.Authors = append(e.Authors, Author{
		ID:   d.ID,
		Name: d.Name,
	})
}

func (e *Book) SetAuthors() {
	if e.m.R.Author != nil {
		e.addAuthor(e.m.R.Author)
	}

	if e.m.R.Author2 != nil {
		e.addAuthor(e.m.R.Author2)
	}

	if e.m.R.Author3 != nil {
		e.addAuthor(e.m.R.Author3)
	}
}

func (e *Book) addTranslator(d *models.Author) {
	e.Translators = append(e.Translators, Author{
		ID:   d.ID,
		Name: d.Name,
	})
}

func (e *Book) SetTranslators() {
	if e.m.R.Translator != nil {
		e.addAuthor(e.m.R.Translator)
	}

	if e.m.R.Translator2 != nil {
		e.addAuthor(e.m.R.Translator2)
	}

	if e.m.R.Translator3 != nil {
		e.addAuthor(e.m.R.Translator3)
	}
}

func (e *Book) addCategory(id int) {
	cat, err := models.BookCategories(qm.Where("id=?", id)).
		OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
		return
	}
	e.Categories = append(e.Categories, Category{
		ID:    cat.ID,
		Title: cat.Title,
	})
}

func (e *Book) SetCategories() {
	for _, c := range e.m.R.BookCategoryAssigns {
		// log.Println(c.CategoryID)
		e.addCategory(c.CategoryID)
	}
}

func (e *Book) SetPublisher() {
	if e.m.R.Publisher == nil {
		return
	}

	e.Publisher = Publisher{
		ID:    e.m.R.Publisher.ID,
		Title: e.m.R.Publisher.Title,
	}
}

var esClient *elastic.Client

func SetEsClient(c *elastic.Client) {
	esClient = c
}

func NewBookByModel(m *models.Book, client *elastic.Client) {
	esBook := Book{
		ID:                 m.ID,
		Title:              m.Title,
		Format:             m.Format,
		ContentType:        m.ContentType,
		Language:           m.Language,
		Price:              m.Price,
		Publish:            m.Publish,
		Free:               m.Free,
		PublishTime:        m.PublishTime.Time,
		DayDownloadCount:   m.R.BookStat.DayDownloadCount,
		WeekDownloadCount:  m.R.BookStat.WeekDownloadCount,
		MonthDownloadCount: m.R.BookStat.MonthDownloadCount,
		AllDownloadCount:   m.R.BookStat.AllDownloadCount,
		DaySalesCount:      m.R.BookStat.DaySalesCount.Int,
		WeekSalesCount:     m.R.BookStat.WeekSalesCount,
		MonthSalesCount:    m.R.BookStat.MonthSalesCount.Int,
		AllSalesCount:      m.R.BookStat.AllSalesCount,
		AllCommentCount:    m.R.BookStat.AllCommentCount.Int,
		m:                  m,
	}

	esBook.SetAuthors()
	esBook.SetTranslators()
	esBook.SetCategories()
	esBook.SetPublisher()

	bookID := fmt.Sprintf("book-%d", m.ID)

	_, err := client.Index().
		Index(ProductIndexName).
		Type(ProductIndexType).
		Id(bookID).
		BodyJson(esBook).
		Do(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	// spew.Dump(res)

}

const (
	ProductIndexName = "fidibo_v3"
	ProductIndexType = "book"
)
