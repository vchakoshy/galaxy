package rest

import (
	// "golang.org/x/tools/go/analysis/passes/unmarshal"
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gitlab.fidibo.com/backend/galaxy/api/models"
)

func AuthorItem(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
	author, err := models.Authors(qm.Where("id=?", id)).
		OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, author)

}

func AuthorItemUpdate(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
	author, err := models.Authors(qm.Where("id=?", id)).
		OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, author)

}

func AuthorItemDelete(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
	author, err := models.Authors(qm.Where("id=?", id)).
		OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, author)

}

func AuthorItemPost(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
	author, err := models.Authors(qm.Where("id=?", id)).
		OneG(context.Background())
	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, author)

}

type sortMap []string

func NewSortMapByGinContext(c *gin.Context) string {
	sm := sortMap{}
	sorts := c.Query("sort")
	json.Unmarshal([]byte(sorts), &sm)
	sortQueryText := strings.Join(sm, " ")
	return sortQueryText
}

func AuthorList(c *gin.Context) {

	sortQueryText := NewSortMapByGinContext(c)

	authors, err := models.
		Authors(
			qm.Limit(30),
			qm.OrderBy(sortQueryText)).
		AllG(context.Background())

	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, authors)
}
