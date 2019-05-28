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

type flexGenericChildAction struct {
	Type      string            `json:"type"`
	Input     []FlexActionInput `json:"input"`
	ExtraData interface{}       `json:"extraData"`
	Method    string            `json:"method"`
}

type flexGenericBook struct {
	ChildAction flexGenericChildAction `json:"childAction"`
	Title       string                 `json:"title"`
	SubTitle    string                 `json:"subTitle"`
	Icon        string                 `json:"icon"`
	Image       string                 `json:"image"`
	Format      string                 `json:"format"`
	ContentType string                 `json:"content_type"`
	Badge       Badge                  `json:"badge"`
	Action      flexGenericChildAction `json:"action"`
	BookID      string                 `json:"bookId"`
}

func newGenericBookByQuery(queries []qm.QueryMod) []flexGenericBook {

	res := make([]flexGenericBook, 0)

	// boil.DebugMode = true

	books, err := models.Books(queries...).AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	for _, b := range books {
		bookIDStr := strconv.Itoa(b.ID)
		fb := flexGenericBook{
			Title:       b.Title,
			SubTitle:    b.SubTitle.String,
			BookID:      bookIDStr,
			Image:       modext.GetBookNormalImage(b),
			Icon:        modext.GetBookNormalImage(b),
			Format:      b.Format,
			ContentType: b.ContentType,
			Action: flexGenericChildAction{
				Type: "book",
				Input: []FlexActionInput{
					FlexActionInput{
						Key:   "bookId",
						Value: bookIDStr,
					},
					FlexActionInput{
						Key:   "pageName",
						Value: "BOOK_OVERVIEW_PAGE",
					},
				},
				Method: "/book/" + bookIDStr + "/get",
			},
			ChildAction: flexGenericChildAction{
				Type: "book",
				Input: []FlexActionInput{
					FlexActionInput{
						Key:   "bookId",
						Value: bookIDStr,
					},
					FlexActionInput{
						Key:   "pageName",
						Value: "BOOK_OVERVIEW_PAGE",
					},
				},
				Method: "/book/" + bookIDStr + "/get",
			},
		}
		res = append(res, fb)
	}

	return res
}
