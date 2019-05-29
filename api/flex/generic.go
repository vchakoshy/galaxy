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

type Generic struct {
	ChildAction *BaseAction `json:"childAction,omitempty"`
	Title       string      `json:"title,omitempty"`
	SubTitle    string      `json:"subTitle,omitempty"`
	Icon        string      `json:"icon,omitempty"`
	IconTitle   string      `json:"iconTitle,omitempty"`
	FooterText  string      `json:"footerText,omitempty"`
	Image       string      `json:"image,omitempty"`
	Ratio       string      `json:"ratio,omitempty"`
	Format      string      `json:"format,omitempty"`
	ContentType string      `json:"content_type,omitempty"`
	Badge       *Badge      `json:"badge,omitempty"`
	Action      *BaseAction `json:"action,omitempty"`
	BookID      string      `json:"bookId,omitempty"`
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
		bookIDStr := strconv.Itoa(b.ID)
		fb := Generic{
			Title:       b.Title,
			SubTitle:    b.SubTitle.String,
			BookID:      bookIDStr,
			Image:       modext.GetBookNormalImage(b),
			Icon:        modext.GetBookNormalImage(b),
			Format:      b.Format,
			ContentType: b.ContentType,
			Action: &BaseAction{
				Type: "book",
				Input: []ComponentAction{
					ComponentAction{
						Key:   "bookId",
						Value: bookIDStr,
					},
					ComponentAction{
						Key:   "pageName",
						Value: "BOOK_OVERVIEW_PAGE",
					},
				},
				Method: "/book/" + bookIDStr + "/get",
			},
			ChildAction: &BaseAction{
				Type: "book",
				Input: []ComponentAction{
					ComponentAction{
						Key:   "bookId",
						Value: bookIDStr,
					},
					ComponentAction{
						Key:   "pageName",
						Value: "BOOK_OVERVIEW_PAGE",
					},
				},
				Method: "/book/" + bookIDStr + "/get",
			},
		}
		res = append(res, fb)

		rs := newBookByModel(b)

		resBook = append(resBook, rs)
	}

	return res, resBook
}
