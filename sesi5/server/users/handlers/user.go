package handlers

import (
	"fmt"
	"net/http"
	"sesi5/server/users/params"
	"sesi5/server/users/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	svc *services.UserServices
}

func NewUserHandler(svc *services.UserServices) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) CreateNewUser(c echo.Context) error {
	traceId := uuid.New().String()
	address := params.Address{}
	var req params.UserCreate = params.UserCreate{
		Name:    "",
		Email:   "",
		Age:     1,
		Address: address,
	}

	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Printf("%+v\n", req)

	resp := u.svc.CreateUser(c.Request().Context(), &req)
	if resp != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("X-trace-id", traceId)
	return c.JSON(http.StatusCreated, M{
		"status":  http.StatusCreated,
		"message": "CREATED SUCCESS",
	})

}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	resp := u.svc.GetUsers(c.Request().Context())
	return c.JSON(http.StatusOK, M{
		"status":  http.StatusOK,
		"message": "GET ALL SUCCESS",
		"payload": resp,
	})

}
