package main

import (
	"sesi5/server/users/handlers"

	pkgError "sesi5/pkg/error"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// example of migrate to microservices
func main() {
	router := echo.New()
	customError := pkgError.CustomValidator{
		Validator: validator.New(),
	}
	router.Validator = &customError
	handlers.BuildUserHandler(router)

	err := router.Start(":6666")
	router.Logger.Fatalf("Fail to running server with error : %v", err.Error())
}
