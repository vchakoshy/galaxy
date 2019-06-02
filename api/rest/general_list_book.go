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
			Setup: flex.Setup{},
		},
	}

	if len(rq.CategoryIds) > 0 {
		cs.Settings.Setup.Category = flex.ProviderSetup{
			Ids: rq.CategoryIds,
		}
	}

	bp := flex.BookDataProvider{
		ComponentSettings: cs,
		Type:              "book",
	}
	out := bp.GetOutputComponent()

	fr := flex.NewResponse()
	fr.Output.Items = out.Data.Items.Model
	fr.Output.Title = "not set "
	fr.Output.Result = true
	fr.Output.IsLastPage = false

	c.JSON(200, fr)

}
