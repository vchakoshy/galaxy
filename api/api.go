package api

import (
	"github.com/gin-gonic/gin"
	"gitlab.fidibo.com/backend/galaxy/api/rest"
)

// Run runs api
func Run() {

	r := gin.Default()

	r.GET("/api/v1/author/", rest.AuthorItem)

	r.Run("0.0.0.0:80")
}
