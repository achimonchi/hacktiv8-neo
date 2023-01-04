package handler

import (
	"fmt"
	"net/http"
	"sesi9/application/model"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

type AuthHandler struct {
	store *sessions.CookieStore
}

func NewAuthHandler(store *sessions.CookieStore) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

type M map[string]interface{}

const (
	CSRF_TokenHeader = "X-CSRF-TOKEN"
	CSRF_Key         = "csrf"
	SESSION_ID       = "SESSIONID"
)

var tmpl = template.Must(template.ParseGlob("./static/*.html"))

func (a *AuthHandler) FormRegister(ctx echo.Context) error {
	data := make(M)
	data[CSRF_Key] = ctx.Get(CSRF_Key)
	return tmpl.ExecuteTemplate(ctx.Response(), "register.html", data)
}

func (a *AuthHandler) FormLogin(ctx echo.Context) error {
	data := make(M)
	data[CSRF_Key] = ctx.Get(CSRF_Key)
	return tmpl.ExecuteTemplate(ctx.Response(), "login.html", data)
}

func (a *AuthHandler) Login(ctx echo.Context) error {
	req := new(model.Auth)
	err := ctx.Bind(&req)
	fmt.Println(err)
	if err != nil {
		return err
	}

	var data = &model.Auth{}
	data = nil
	for _, auth := range model.AuthData {
		if auth.Email == req.Email && auth.Password == req.Password {
			data = &auth
			break
		}
	}

	if data == nil {
		return ctx.JSON(http.StatusNotFound, M{
			"error": "not found",
		})
	}

	payload := M{
		"payload": data,
		"status":  http.StatusOK,
	}

	sess, _ := a.store.Get(ctx.Request(), SESSION_ID)

	sess.Values["is_login"] = true
	sess.Values["data"] = data.Email

	err = sess.Save(ctx.Request(), ctx.Response())
	fmt.Println("try to save session :", err)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, payload)
}
func (a *AuthHandler) Register(ctx echo.Context) error {
	var req = new(model.Auth)

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	req.Id = len(model.AuthData) + 1

	model.AuthData = append(model.AuthData, *req)
	return ctx.JSON(http.StatusCreated, M{
		"message": "register success",
		"status":  http.StatusCreated,
	})
}

func (a *AuthHandler) GetCSRF(ctx echo.Context) error {
	data := make(M)
	data[CSRF_Key] = ctx.Get(CSRF_Key)

	return ctx.JSON(http.StatusOK, data)
}

func (a *AuthHandler) GetAll(ctx echo.Context) error {
	session, _ := a.store.Get(ctx.Request(), SESSION_ID)

	// validate session is no data
	if len(session.Values) == 0 {
		return ctx.JSON(http.StatusUnauthorized, M{
			"status": "fail",
			"error":  "no sessions",
		})
	}

	return ctx.JSON(http.StatusOK, M{
		"user": M{
			"email": session.Values["data"],
		},
		"payload": model.AuthData,
	})
}
