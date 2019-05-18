package api

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	r.Use(corsConfig())

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/authors/item/:id/", rest.AuthorItem)
		v1.PUT("/authors/item/:id/", rest.AuthorItemUpdate)
		v1.POST("/authors/item/:id/", rest.AuthorItemPost)
		v1.DELETE("/authors/item/:id/", rest.AuthorItemDelete)

		v1.GET("/authors/list/", rest.AuthorList)

		v1.POST("/flex/page/blank", rest.PageBlank)
	}

	err = r.Run("0.0.0.0:8080")
	log.Println(err.Error())
}

func corsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	return cors.New(config)
}
