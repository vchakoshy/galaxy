package rest

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"gitlab.fidibo.com/backend/galaxy/api/flex"
)

func GeneralList(c *gin.Context) {
	provider := c.Param("provider")
	rq, err := NewPageReqDataFromRequestBody(c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{"error": "bad data"})
		return
	}
	spew.Dump(rq)

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
	cs.Settings.Setup.Sort.Type = "sort"
	cs.Settings.Setup.Sort.Value = rq.Sort

	out := flex.NewDataProviderByComponentSettings(cs, provider).GetOutputComponent()

	fr := flex.NewResponse()
	fr.Output.Items = out.Data.Items.Model
	fr.Output.Title = "not set "
	fr.Output.Result = true
	fr.Output.IsLastPage = false

	c.JSON(200, fr)

}
