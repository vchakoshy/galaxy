package rest

import (
	"github.com/gin-gonic/gin"
)

func AuthorItem(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
