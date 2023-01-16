package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
