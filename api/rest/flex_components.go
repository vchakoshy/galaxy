package rest

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

func PageBlank(c *gin.Context) {
	pageID := c.Param("id")

	fp, err := models.FlexPages(qm.Where("id=?", pageID)).
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
					Key:   "pageID",
					Value: tabs.PageID,
				},
			},
		}
	}

	fs := flexStruct{
		Output: flexOutput{
			Components: []string{},
			FlexErrors: []string{},
			Setting:    settings,
			Title:      fp.Title,
			Result:     true,
			IsLastPage: true,
		},
	}

	c.JSON(200, fs)

}
