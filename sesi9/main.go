package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	CSRF_TokenHeader = "X-CSRF-TOKEN"
	CSRF_Key         = "csrf"
)

type M map[string]interface{}

func main() {
	tmpl := template.Must(template.ParseGlob("./static/*.html"))

	e := echo.New()

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRF_TokenHeader,
		ContextKey:  CSRF_Key,
	}))

	e.GET("/index", func(ctx echo.Context) error {
		data := make(M)
		data[CSRF_Key] = ctx.Get(CSRF_Key)
		return tmpl.Execute(ctx.Response(), data)
	})

	e.POST("/sayhello", func(ctx echo.Context) error {
		data := make(M)

		err := ctx.Bind(&data)
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Hello, %s", data["name"])
		data["message"] = message
		return ctx.JSON(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":4444"))
}
