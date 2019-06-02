package rest

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type pageReqData struct {
	PageName string `json:"pageName"`
	PageID   string `json:"pageId"`
	Page     int    `json:"page"`
}

func NewPageReqDataFromRequestBody(c *gin.Context) (d pageReqData, err error) {
	err = json.NewDecoder(c.Request.Body).Decode(&d)

	// fix default zero page number
	if d.Page == 0 {
		d.Page = 1
	}

	return
}

func (p pageReqData) getPageQuery() (q qm.QueryMod) {
	if p.PageID != "" {
		pageIDInt, _ := strconv.Atoi(p.PageID)
		q = qm.Where("id=?", pageIDInt)
	} else {
		q = qm.Where("name=?", p.PageName)
	}
	return q
}
