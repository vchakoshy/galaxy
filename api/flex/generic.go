package flex

import (
	"fmt"
	"strconv"
	"strings"

	"context"
	"log"

	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
	"gitlab.fidibo.com/backend/galaxy/api/modext"
)

// Generic struct
type Generic struct {
	ChildAction  *Action `json:"childAction,omitempty"`
	Title        string  `json:"title,omitempty"`
	SubTitle     string  `json:"subTitle,omitempty"`
	Icon         string  `json:"icon,omitempty"`
	IconSubtitle string  `json:"iconSubtitle,omitempty"`
	IconTitle    string  `json:"iconTitle,omitempty"`
	FooterText   string  `json:"footerText,omitempty"`
	Image        string  `json:"image,omitempty"`
	Ratio        string  `json:"ratio,omitempty"`
	Format       string  `json:"format,omitempty"`
	ContentType  string  `json:"content_type,omitempty"`
	Badge        *Badge  `json:"badge,omitempty"`
	Action       *Action `json:"action,omitempty"`
	BookID       string  `json:"bookId,omitempty"`
	Rate         string  `json:"rate,omitempty"`
}

func newBooksByIds(ids []int) (gens []Generic, books []Book) {
	gens = make([]Generic, 0)
	books = make([]Book, 0)

	iids := make([]interface{}, len(ids))
	idStrs := make([]string, len(ids))
	for index, num := range ids {
		iids[index] = num
		idStrs[index] = strconv.Itoa(num)
	}

	orderByIdid := fmt.Sprintf("FIELD(id,%s)", strings.Join(idStrs, ","))

	bs, err := models.Books(
		qm.WhereIn("id in ?", iids...),
		qm.OrderBy(orderByIdid),
		qm.Load("Publisher"),
		qm.Load("Author")).
		AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, b := range bs {
		gen := newGenericByModel(b)
		book := newBookByModel(b)
		gens = append(gens, gen)
		books = append(books, book)
	}

	return
}

func newGenericByModel(b *models.Book) Generic {
	bookIDStr := strconv.Itoa(b.ID)
	fb := Generic{
		Title:       b.Title,
		BookID:      bookIDStr,
		Image:       modext.GetBookNormalImage(b),
		Icon:        modext.GetBookNormalImage(b),
		Format:      strings.ToLower(b.Format),
		ContentType: strings.ToUpper(b.ContentType),
		Rate:        fmt.Sprintf("%.2f", b.Rate),
		Action: &Action{
			Type: "book",
			Input: []InputAction{
				InputAction{
					Key:   "bookId",
					Value: bookIDStr,
				},
				InputAction{
					Key:   "pageName",
					Value: "BOOK_OVERVIEW_PAGE",
				},
			},
			Method: "/book/" + bookIDStr + "/get",
		},
		ChildAction: &Action{
			Type: "book",
			Input: []InputAction{
				InputAction{
					Key:   "bookId",
					Value: bookIDStr,
				},
				InputAction{
					Key:   "pageName",
					Value: "BOOK_OVERVIEW_PAGE",
				},
			},
			Method: "/book/" + bookIDStr + "/get",
		},
	}

	if b.R.Author != nil {
		fb.SubTitle = b.R.Author.Name
	}

	return fb
}
