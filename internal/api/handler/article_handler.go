package handler

import (
	"feelgoodfeatures/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {

	articles, _ := service.ScrapeNews()

	c.JSON(http.StatusOK, articles)
}
