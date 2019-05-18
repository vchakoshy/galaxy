package rest

import (
	"context"
	"log"

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

func AuthorList(c *gin.Context) {
	authors, err := models.
		Authors(
			qm.Limit(30),
			qm.OrderBy("id desc")).
		AllG(context.Background())

	if err != nil {
		log.Println(err.Error())
	}
	c.JSON(200, authors)
}
