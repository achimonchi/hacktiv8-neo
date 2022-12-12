package main

import (
	"fmt"
	"net/http"
	"sesi5/server/users/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(gin.Recovery())

	v1 := router.Group("/v1")
	v2 := router.Group("/v2")
	userGroup := v1.Group("/users")
	{
		router.Use(Log)
		userGroup.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, handlers.M{
				"status":  http.StatusOK,
				"message": "healthy",
			})
		})

		router.Use(CheckAdmin)
		userGroup.POST("", func(ctx *gin.Context) {

		})

		userGroup.PUT("/:id", func(ctx *gin.Context) {

		})
	}

	productGroup := v2.Group("/products")
	{
		productGroup.Use(Log)
		productGroup.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, handlers.M{
				"status":  http.StatusOK,
				"message": "healthy",
			})
		})

		productGroup.POST("", Log, Auth, CheckIsUser, func(ctx *gin.Context) {

		})

		productGroup.PUT("/:id", func(ctx *gin.Context) {

		})
	}

	router.Run(":7777")
}

func Log(ctx *gin.Context) {
	fmt.Println("request method :", ctx.Request.Method)
	ctx.Next()
}
