package handlers

import (
	"sesi5/server/users/models"
	"sesi5/server/users/repositories"
	"sesi5/server/users/services"

	"github.com/labstack/echo/v4"
)

func BuildUserHandler(router *echo.Echo) {
	repo := repositories.NewUserRepo(&models.Users)
	svc := services.NewUserServices(repo)
	controller := NewUserHandler(svc)

	// generate router
	router.POST("/users", controller.CreateNewUser)
	router.GET("/users", controller.GetAllUsers)
}
