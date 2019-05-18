package rest

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

type reqData struct {
	PageName string `json:"pageName"`
}

func PageBlank(c *gin.Context) {
	pageID := c.Param("id")

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
	if pageID != "" {
		q = qm.Where("id=?", pageID)
	} else {
		q = qm.Where("name=?", rd.PageName)
	}
	log.Println(q)

	fp, err := models.FlexPages(q).
		OneG(context.Background())

	if err != nil {
		log.Println(err.Error())
	}

	settings := flexSetting{}

	err = json.Unmarshal([]byte(fp.SettingData), &settings)
	if err != nil {
		log.Println(err.Error())
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
			Components: []flexComponent{},
			FlexErrors: []string{},
			Setting:    settings,
			Title:      fp.Title,
			Result:     true,
			IsLastPage: true,
		},
	}

	c.JSON(200, fs)

}
