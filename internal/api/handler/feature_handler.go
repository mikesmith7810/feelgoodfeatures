package handler

import (
	"net/http"

	"feelgoodfeatures/internal/domain"
	"github.com/gin-gonic/gin"
)

func GetFeature(c *gin.Context) {
	feature := domain.Feature{
		ID:      "1",
		Content: "This is a nice story about a cat.",
	}

	c.JSON(http.StatusOK, feature)
}
