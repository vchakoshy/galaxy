package api

import (
	"log"

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

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/authors/item/:id/", rest.AuthorItem)
		v1.PUT("/authors/item/:id/", rest.AuthorItemUpdate)
		v1.POST("/authors/item/:id/", rest.AuthorItemPost)
		v1.DELETE("/authors/item/:id/", rest.AuthorItemDelete)

		v1.GET("/authors/list/", rest.AuthorList)
		v1.OPTIONS("/authors/list/", rest.AuthorList)

		v1.POST("/flex/page/blank", rest.PageBlank)
	}

	go func() {
		hubble.Run()
	}()

	err = r.Run("0.0.0.0:8080")
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
