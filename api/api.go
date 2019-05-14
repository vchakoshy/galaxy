package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.fidibo.com/backend/galaxy/api/model"
	"gitlab.fidibo.com/backend/galaxy/api/rest"
)

// Run runs api
func Run() {

	mysqlDNS := "reader:j3AhwCj4SqVQI62Y@tcp(79.175.173.69:3306)/fidibo1_fidibo?autocommit=true"
	db, err := sqlx.Open("mysql", mysqlDNS)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer db.Close()

	model.SetDB(db)

	r := gin.Default()

	r.GET("/api/v1/author/", rest.AuthorItem)

	r.Run("0.0.0.0:80")
}
