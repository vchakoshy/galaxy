package rest 

import (
	"gitlab.fidibo.com/backend/galaxy/api/modext"
	
	"strconv"
	"log"
	"context"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)


type flexGenericChildAction struct {
	Type      string            `json:"type"`
	Input     []flexActionInput `json:"input"`
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
	Badge       struct {
		Text     string `json:"text"`
		BgColor  string `json:"bg_color"`
		TxtColor string `json:"txt_color"`
	} `json:"badge"`
	Action flexGenericChildAction `json:"action"`
	BookID string                 `json:"bookId"`
}

func newGenericBookByID(bookID string) flexGenericBook {
	b, err := models.Books(qm.Where("id=?", bookID)).OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
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
			Input: []flexActionInput{
				flexActionInput{
					Key:   "bookId",
					Value: bookIDStr,
				},
				flexActionInput{
					Key:   "pageName",
					Value: "BOOK_OVERVIEW_PAGE",
				},
			},
			Method: "/book/" + bookIDStr + "/get",
		},
		ChildAction: flexGenericChildAction{
			Type: "book",
			Input: []flexActionInput{
				flexActionInput{
					Key:   "bookId",
					Value: bookIDStr,
				},
				flexActionInput{
					Key:   "pageName",
					Value: "BOOK_OVERVIEW_PAGE",
				},
			},
			Method: "/book/" + bookIDStr + "/get",
		},
	}
	return fb 
}