package rest

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.fidibo.com/backend/galaxy/api/flex"
)

func GeneralListBook(c *gin.Context) {
	rq, err := NewPageReqDataFromRequestBody(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "bad data"})
		return
	}

	cs := flex.ComponentSettings{
		Settings: flex.Settings{
			Setup: flex.Setup{
				Category: flex.ProviderSetup{
					Ids: rq.CategoryIds,
				},
			},
		},
	}

	bp := flex.BookDataProvider{
		ComponentSettings: cs,
		Type:              "book",
	}
	out := bp.GetOutputComponent()

	fs := flex.Response{
		Output: flex.Output{
			Items:      out.Data.Items.Model,
			FlexErrors: []string{},
			Title:      "not set ",
			Result:     true,
			IsLastPage: false,
		},
	}

	c.JSON(200, fs)

}
