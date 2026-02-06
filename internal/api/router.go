package api

import (
	"feelgoodfeatures/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/feature", handler.GetArticles)
	}

	return r
}
