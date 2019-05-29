package flex

import (
	"strconv"
	"time"

	"context"
	"log"

	// "github.com/volatiletech/sqlboiler/boil"
	"github.com/patrickmn/go-cache"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
	"gitlab.fidibo.com/backend/galaxy/api/modext"
)

type GenericChildAction struct {
	Type      string        `json:"type"`
	Input     []ActionInput `json:"input"`
	ExtraData interface{}   `json:"extraData"`
	Method    string        `json:"method"`
}

type Generic struct {
	ChildAction GenericChildAction `json:"childAction"`
	Title       string             `json:"title"`
	SubTitle    string             `json:"subTitle"`
	Icon        string             `json:"icon"`
	Image       string             `json:"image"`
	Format      string             `json:"format"`
	ContentType string             `json:"content_type"`
	Badge       Badge              `json:"badge"`
	Action      GenericChildAction `json:"action"`
	BookID      string             `json:"bookId,omitempty"`
}

var genericsCache = cache.New(time.Minute*5, time.Minute*10)

type genericCache struct {
	Generic Generic
	Book    Book
}

func newGenericBookByID(id int) (Generic, Book) {
	cacheKey := strconv.Itoa(id)
	if val, ex := genericsCache.Get(cacheKey); ex {
		s := val.(genericCache)
		return s.Generic, s.Book
	}
	boil.DebugMode = true

	book, err := models.
		Books(
			qm.Where("id=?", id),
			qm.Load("Publisher"),
			qm.Load("Author")).
		OneG(context.Background())

	if err != nil {
		return Generic{}, Book{}
	}

	fb := newGenericByModel(book)
	rs := newBookByModel(book)
	genericsCache.Set(cacheKey, genericCache{fb, rs}, cache.DefaultExpiration)
	return fb, rs
}

func newGenericBookByQuery(queries []qm.QueryMod) ([]Generic, []Book) {
	res := make([]Generic, 0)
	resBook := make([]Book, 0)

	queries = append(queries, qm.Load("Publisher"), qm.Load("Author"))

	books, err := models.Books(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	for _, b := range books {

		fb := newGenericByModel(b)
		res = append(res, fb)

		rs := newBookByModel(b)
		resBook = append(resBook, rs)
	}

	return res, resBook
}

func newGenericBookByIds(ids []int) ([]Generic, []Book) {
	gens := make([]Generic, 0)
	books := make([]Book, 0)

	for _, id := range ids {
		gen, book := newGenericBookByID(id)
		gens = append(gens, gen)
		books = append(books, book)

	}

	return gens, books
}

func newGenericByModel(b *models.Book) Generic {
	bookIDStr := strconv.Itoa(b.ID)
	fb := Generic{
		Title:       b.Title,
		SubTitle:    b.SubTitle.String,
		BookID:      bookIDStr,
		Image:       modext.GetBookNormalImage(b),
		Icon:        modext.GetBookNormalImage(b),
		Format:      b.Format,
		ContentType: b.ContentType,
		Action: GenericChildAction{
			Type: "book",
			Input: []ActionInput{
				ActionInput{
					Key:   "bookId",
					Value: bookIDStr,
				},
				ActionInput{
					Key:   "pageName",
					Value: "BOOK_OVERVIEW_PAGE",
				},
			},
			Method: "/book/" + bookIDStr + "/get",
		},
		ChildAction: GenericChildAction{
			Type: "book",
			Input: []ActionInput{
				ActionInput{
					Key:   "bookId",
					Value: bookIDStr,
				},
				ActionInput{
					Key:   "pageName",
					Value: "BOOK_OVERVIEW_PAGE",
				},
			},
			Method: "/book/" + bookIDStr + "/get",
		},
	}

	return fb
}
