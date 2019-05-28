package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	req, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(req))

	err = json.Unmarshal(req, &d)
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
