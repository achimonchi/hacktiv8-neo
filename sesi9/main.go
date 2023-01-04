package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sesi9/application/handler"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
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
	e.GET("/view", func(ctx echo.Context) error {
		data := make(M)
		data[CSRF_Key] = ctx.Get(CSRF_Key)
		return tmpl.ExecuteTemplate(ctx.Response(), "view.html", data)
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

	key := securecookie.GenerateRandomKey(32)
	store := sessions.NewCookieStore(key)

	handler := handler.NewAuthHandler(store)

	e.POST("/auth/register", handler.Register)
	e.POST("/auth/login", handler.Login)
	e.GET("/form/register", handler.FormRegister)
	e.GET("/form/login", handler.FormLogin)
	e.GET("/auth", handler.GetAll, CheckClientId())
	e.GET("/auth/csrfkey", handler.GetCSRF)

	e.Logger.Fatal(e.Start(":4444"))
}

func CheckClientId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			header := ctx.Request().Header.Get("X-CLIENT-ID")
			fmt.Println(header)
			if header != "backend" {
				return ctx.JSON(http.StatusForbidden, M{
					"message": "forbidden access",
				})
			}
			return next(ctx)
		}
	}
}
