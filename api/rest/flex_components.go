package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

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

		cs := flexComponentSettings{}
		json.Unmarshal([]byte(comp.ComponentSetting.String), &cs)

		log.Println(cs.Settings.DataProvider)

		if cs.Settings.DataProvider == "BOOK" {
			set := cs.Settings.Setup.Book

			queries := make([]qm.QueryMod, 0)

			qidis := make([]interface{}, len(set.Ids))
			for _, id := range set.Ids {
				qidis = append(qidis, id)
			}

			if len(qidis) > 0 {
				queries = append(queries, qm.WhereIn("id in ?", qidis...))
			}

			queries = append(queries, qm.Limit(30))

			// if set.Type == "STATIC" {
			com := flexComponent{
				Type:         "HL_BOOKS_ARTICLE",
				ResourceType: "BOOK",
			}

			log.Println(queries)

			res := newGenericBookByQuery(queries)

			com.Data.Items.Generic = make([]interface{}, len(res))
			for i, v := range res {
				com.Data.Items.Generic[i] = v
			}

			componenets = append(componenets, com)
			// }

		}

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
