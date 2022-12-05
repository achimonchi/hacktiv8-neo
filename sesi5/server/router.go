package server

import (
	pkgError "sesi5/pkg/error"
	"sesi5/server/users/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Router struct {
	router *echo.Echo
	port   string
}

func NewRouter(port string) *Router {
	router := echo.New()
	customError := pkgError.CustomValidator{
		Validator: validator.New(),
	}
	router.Validator = &customError
	return &Router{
		router: router,
		port:   port,
	}
}

func (r *Router) Start() {
	handlers.BuildUserHandler(r.router)

	err := r.router.Start(r.port)
	r.router.Logger.Fatalf("Fail to running server with error : %v", err.Error())
}
