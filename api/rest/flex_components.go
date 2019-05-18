package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

type reqData struct {
	PageName string `json:"pageName"`
	PageID   string `json:"pageId"`
}

func PageBlank(c *gin.Context) {
	req, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
	}

	var rd reqData
	err = json.Unmarshal(req, &rd)
	if err != nil {
		log.Println(err.Error())
	}

	var q qm.QueryMod
	if rd.PageID != "" {
		pageIDInt, _ := strconv.Atoi(rd.PageID)
		q = qm.Where("id=?", pageIDInt)
	} else {
		q = qm.Where("name=?", rd.PageName)
	}
	log.Println(q)

	fp, err := models.FlexPages(q).OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	qComponent := qm.Where("page_id=?", fp.ID)
	log.Println(qComponent)
	fpc, err := models.FlexPageComponents(qComponent, qm.OrderBy("crud_order asc")).
		AllG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}

	settings := flexSetting{}

	err = json.Unmarshal([]byte(fp.SettingData), &settings)
	if err != nil {
		log.Println(err.Error())
	}

	componenets := make([]flexComponent, 0)

	for _, comp := range fpc {
		com := flexComponent{
			Type:         "HL_BOOKS_ARTICLE",
			ResourceType: "BOOK",
		}

		cs := flexComponentSettings{}
		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)
		if cs.Settings.DataProvider == "BOOK" {
			set := cs.Settings.Setup.Book
			log.Println("data provider is book")
			// spew.Dump(set)
			if set.Type == "STATIC" && set.Mapping == "id" {
				idis := set.Ids
				for _, id := range idis {
					b, err := models.Books(qm.Where("id=?", id)).OneG(context.Background())
					if err != nil {
						log.Println(err.Error())
					}
					bookIDStr := strconv.Itoa(b.ID)
					fb := flexGenericBook{
						Title:       b.Title,
						SubTitle:    b.SubTitle.String,
						BookID:      bookIDStr,
						Image:       "/images/books/" + strings.Replace(b.ImageName.String, ".jpg", "_normal.jpg", 1),
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

					com.Data.Items.Generic = append(com.Data.Items.Generic, fb)

				}
			}

		}
		componenets = append(componenets, com)
	}

	for index, tabs := range settings.Tabs {
		settings.Tabs[index].Action = flexAction{
			Method: "/flex/page/blank",
			Type:   "flex_blank_page",
			Input: []flexActionInput{
				flexActionInput{
					Key:   "pageId",
					Value: tabs.PageID,
				},
			},
		}
	}

	fs := flexStruct{
		Output: flexOutput{
			Components: componenets,
			FlexErrors: []string{},
			Setting:    settings,
			Title:      fp.Title,
			Result:     true,
			IsLastPage: true,
		},
	}

	c.JSON(200, fs)

}
