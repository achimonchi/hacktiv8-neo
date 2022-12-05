package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	return c.validator.Struct(i)
}

type M map[string]interface{}

func main() {
	router := echo.New()
	router.Validator = &CustomValidator{validator: validator.New()}

	router.HTTPErrorHandler = func(err error, ctx echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		ctx.Logger().Error(report)
		ctx.JSON(report.Code, report)
	}

	router.GET("/", HandleIndex)
	router.GET("/users", HandleUser)
	router.GET("/products/:productId/*", func(c echo.Context) error {
		productId := c.Param("productId")
		moreParams := c.Param("*")
		data := fmt.Sprintf("Try to get product with id : %s and more params is : %s", productId, moreParams)

		return c.JSON(http.StatusOK, M{
			"status": http.StatusOK,
			"data":   data,
		})
	})

	router.POST("/users", func(c echo.Context) error {
		type req struct {
			Name  string `validate:"required"`
			Age   int    `validate:"required,gte=1"`
			Grade int    `validate:"required,gte=0,lte=6"`
		}

		data := req{}

		err := c.Bind(&data)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, M{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
		}

		err = c.Validate(data)
		if err != nil {
			errorData := map[string]interface{}{}
			for _, err := range err.(validator.ValidationErrors) {
				errorData[strings.ToLower(err.Field())] = err.Tag()
				// errorData += fmt.Sprintf("field %v should %v param %v and value %v", err.Field(), err.Tag(), err.Param(), err.Value())
			}
			// c.JSON(http.StatusBadRequest)
			return echo.NewHTTPError(http.StatusBadRequest, M{
				"status":          http.StatusBadRequest,
				"error":           "BAD_REQUEST_VALIDATION",
				"additional_info": errorData,
			})
		}

		return c.JSON(http.StatusOK, M{
			"status": http.StatusOK,
			"data":   data,
		})
	})

	router.GET("/transactions", echo.WrapHandler(
		http.HandlerFunc(HandleTransactions),
	))

	router.Static("/static", "./static")

	err := router.Start(":5555")
	router.Logger.Fatalf("Fail to running server with error : %v", err.Error())
}

func HandleIndex(c echo.Context) error {
	data := M{
		"status": "ok",
		"health": "good",
	}
	return c.JSON(http.StatusOK, data)
}

func HandleUser(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, M{
			"status": http.StatusBadRequest,
			"error":  "name in query is required",
		})
	}

	data := fmt.Sprintf("hello, %s", name)

	return c.JSON(http.StatusOK, M{
		"status": http.StatusOK,
		"data":   data,
	})
}

func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handle Transactions"))
}
