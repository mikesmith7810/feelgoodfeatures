package main

import (
	"feelgoodfeatures/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define Routes
	r.GET("/feature", handler.GetFeature)

	// Run the server on port 8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
