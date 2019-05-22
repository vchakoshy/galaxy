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
		settings.Tabs[index].Action = flex.FlexAction{
			Method: "/flex/page/blank",
			Type:   "flex_blank_page",
			Input: []flex.FlexActionInput{
				flex.FlexActionInput{
					Key:   "pageId",
					Value: tabs.PageID,
				},
			},
		}
	}

	fs := flex.FlexStruct{
		Output: flex.FlexOutput{
			Components: flex.NewComponentByPage(fp.ID).GetData(),
			FlexErrors: []string{},
			Setting:    settings,
			Title:      fp.Title,
			Result:     true,
			IsLastPage: true,
		},
	}

	c.JSON(200, fs)

}
