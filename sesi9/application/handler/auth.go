package handler

import "github.com/labstack/echo"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a *AuthHandler) Register(ctx echo.Context) error {
	panic("")
}

func (a *AuthHandler) Login(ctx echo.Context) error {
	panic("")
}

func (a *AuthHandler) GetAll(ctx echo.Context) error {
	panic("")
}
