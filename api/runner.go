package api 

import (
	"log"
	"github.com/gin-gonic/gin"
)

// Run runs api
func Run(){
	log.Println("api runner start")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

