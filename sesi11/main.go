package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	// db, err := pkg.ConnectDB()
	// if err != nil {
	// 	panic(err)
	// }

	// if db != nil {
	// 	fmt.Println("db connected")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
