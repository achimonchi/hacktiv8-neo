package main

import (
	"fmt"
	"net/http"
	"sesi6/db"

	"github.com/gorilla/context"
	"github.com/labstack/echo/v4"
)

const SESSION_ID = "SESSIONID"

type M map[string]interface{}

func main() {
	e := echo.New()

	// proses pembuatan key
	// expect jangan diletakin di hardcode
	// prefer gunain env file / file config lainnya
	// key := securecookie.GenerateRandomKey(32)
	// proses pembuatan session cookiestore
	// store := sessions.NewCookieStore(key)
	// dbPostgres, _ := db.NewPostgres()

	// store, _ := db.NewPostgresStore(dbPostgres)
	store, _ := db.NewRedisStore()
	store.Options.MaxAge = 10

	// should be used
	// untuk menghindari memori leak
	e.Use(echo.WrapMiddleware(context.ClearHandler))

	e.GET("/set", func(c echo.Context) error {
		query := c.QueryParam("name")
		// get session by session_id
		session, _ := store.Get(c.Request(), SESSION_ID)
		session.Values["name"] = query
		session.Values["age"] = "24"

		// set expire time for session
		// in seconds
		// session.Options.MaxAge = 10
		// session.Options.Secure = true
		// session.Options.HttpOnly = true

		// store session
		err := session.Save(c.Request(), c.Response())
		if err != nil {
			fmt.Println("Error", err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, M{
			"status":  "OK",
			"message": "store cookie success",
		})
	})

	e.GET("/get", func(c echo.Context) error {
		// process get session
		session, _ := store.Get(c.Request(), SESSION_ID)

		// validate session is no data
		if len(session.Values) == 0 {
			return c.JSON(http.StatusOK, M{
				"status": "fail",
				"error":  "no sessions",
			})
		}

		return c.JSON(http.StatusOK, M{
			"data": M{
				"name": session.Values["name"],
				"age":  session.Values["age"],
			},
		})
	})

	e.GET("/delete", func(c echo.Context) error {
		// process get session
		session, _ := store.Get(c.Request(), SESSION_ID)

		// process to expired session
		session.Options.MaxAge = -1
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusTemporaryRedirect, "/get")

	})

	e.Start(":4444")
}
