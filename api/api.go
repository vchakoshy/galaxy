package api

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"

	"gitlab.fidibo.com/backend/galaxy/api/rest"
	"gitlab.fidibo.com/backend/galaxy/hubble"
)

// Run runs api
func Run() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mysqlDNS := "reader:j3AhwCj4SqVQI62Y@tcp(79.175.173.69:3306)/fidibo1_fidibo?autocommit=true&parseTime=true"
	db, err := sqlx.Open("mysql", mysqlDNS)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	boil.SetDB(db)

	r := gin.Default()
	r.Use(CORSMiddleware())

	ping := r.Group("/ping/")
	{
		ping.GET("/pong/", rest.Ping)
		ping.GET("/http/", rest.PingHttp)
	}

	v1 := r.Group("/api/v1/")
	{
		authors := v1.Group("/authors/")
		{
			authors.GET("/item/:id/", rest.AuthorItem)
			authors.PUT("/item/:id/", rest.AuthorItemUpdate)
			authors.POST("/item/:id/", rest.AuthorItemPost)
			authors.DELETE("/item/:id/", rest.AuthorItemDelete)
			authors.GET("/list/", rest.AuthorList)
			authors.OPTIONS("/list/", rest.AuthorList)
		}

		flexy := v1.Group("/flex/")
		{
			flexy.POST("/page/blank", rest.PageBlank)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/") {
			rest.PastVersion(c)
			return
		}

		c.JSON(404, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
			"url":     c.Request.URL.Path,
		})
	})

	if os.Getenv("HUBBLE_REINDEXER") == "1" {
		go func() {
			hubble.Run()
		}()
	}

	listenAddr := os.Getenv("LISTEN_ADDRESS")
	if listenAddr == "" {
		listenAddr = "0.0.0.0:8080"
	}

	err = r.Run(listenAddr)
	log.Println(err.Error())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
