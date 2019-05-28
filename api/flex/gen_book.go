package flex

import (
	"strconv"

	"context"
	"log"

	// "github.com/volatiletech/sqlboiler/boil"
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

type GenericBook struct {
	ChildAction GenericChildAction `json:"childAction"`
	Title       string             `json:"title"`
	SubTitle    string             `json:"subTitle"`
	Icon        string             `json:"icon"`
	Image       string             `json:"image"`
	Format      string             `json:"format"`
	ContentType string             `json:"content_type"`
	Badge       Badge              `json:"badge"`
	Action      GenericChildAction `json:"action"`
	BookID      string             `json:"bookId"`
}

func newGenericBookByQuery(queries []qm.QueryMod) ([]GenericBook, []Book) {
	res := make([]GenericBook, 0)
	resBook := make([]Book, 0)

	queries = append(queries, qm.Load("Publisher"), qm.Load("Author"))

	books, err := models.Books(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	for _, b := range books {
		bookIDStr := strconv.Itoa(b.ID)
		fb := GenericBook{
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
		res = append(res, fb)

		rs := Book{
			BookID:          b.ID,
			BookTitle:       b.Title,
			Price:           b.Price,
			PaperPrice:      b.PaperPrice.Float32,
			BookImage:       modext.GetBookNormalImage(b),
			BookImageSquare: nil,
			Author:          b.R.Author.Name,
			AuthorID:        b.R.Author.ID,
			Translator:      nil,
			TranslatorID:    nil,
			Narrator:        nil,
			NarratorID:      nil,
			PublisherID:     b.R.Publisher.ID,
			PublisherTitle:  b.R.Publisher.Title,
			Rtl:             modext.IsRtl(b),
			Format:          b.Format,
			ContentType:     b.ContentType,
			Rate:            b.Rate,
			Free:            b.Free,
			Path:            "",
			PublishDate:     "",
			RateCount:       "",
			Duration:        "",
			FileSize:        "",
			CategoryID:      "",
			CategoryName:    "",
			PageCount:       "",
			Price2:          "",
			Type:            "",
			SampleCrc:       "",
			Pass:            "",
			Crc:             nil,
			LinkTitle:       "",
			LanguageTitle:   "",
			PriceTitle:      "",
			PriceTitle2:     "",
			Description:     "",
			LinkURL:         "",
			Bought:          false,
			New:             false,
			Featured:        false,
			Favorite:        false,
			MyRate:          0,
			AuthorLogo:      "",
		}

		resBook = append(resBook, rs)
	}

	return res, resBook
}
