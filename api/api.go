package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/boil"

	"gitlab.fidibo.com/backend/galaxy/api/rest"
)

// Run runs api
func Run() {

	mysqlDNS := "reader:j3AhwCj4SqVQI62Y@tcp(79.175.173.69:3306)/fidibo1_fidibo?autocommit=true&parseTime=true"
	db, err := sqlx.Open("mysql", mysqlDNS)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	boil.SetDB(db)

	r := gin.Default()

	r.GET("/api/v1/authors/item/:id/", rest.AuthorItem)
	r.GET("/api/v1/authors/list/", rest.AuthorList)

	r.GET("/api/v1/flex/page/blank/:id", rest.PageBlank)
	r.POST("/api/v1/flex/page/blank", rest.PageBlank)

	err = r.Run("0.0.0.0:8080")
	log.Println(err.Error())
}
