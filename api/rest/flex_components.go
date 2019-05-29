package rest

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.fidibo.com/backend/galaxy/api/flex"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

func PageBlank(c *gin.Context) {
	rq, err := NewPageReqDataFromRequestBody(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "bad data"})
		return
	}

	fp, err := models.
		FlexPages(rq.getPageQuery()).
		OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "bad data"})
		return
	}

	settings := flex.Setting{}

	err = json.Unmarshal([]byte(fp.SettingData), &settings)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "bad data"})
		return
	}

	for index, tabs := range settings.Tabs {
		settings.Tabs[index].Action = flex.BaseAction{
			Method: "/flex/page/blank",
			Type:   "flex_blank_page",
			Input: []flex.ComponentAction{
				flex.ComponentAction{
					Key:   "pageId",
					Value: tabs.PageID,
				},
			},
		}
	}

	cData := flex.NewComponentByPage(fp.ID).GetData(rq.Page)

	isLastPage := false
	if len(cData) == 0 {
		isLastPage = true
	}

	fs := flex.Response{
		Output: flex.Output{
			Components: cData,
			FlexErrors: []string{},
			Setting:    settings,
			Title:      fp.Title,
			Result:     true,
			IsLastPage: isLastPage,
		},
	}

	c.JSON(200, fs)

}
