package api

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"

	"gitlab.fidibo.com/backend/galaxy/api/rest"
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

	// r.Use(corsMiddleware())
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

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

	err = r.Run("0.0.0.0:8080")
	log.Println(err.Error())
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
