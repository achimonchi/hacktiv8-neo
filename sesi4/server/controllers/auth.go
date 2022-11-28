package controllers

import (
	"encoding/json"
	"net/http"
	"sesi4/helper"
	"text/template"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) HandleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	ok, data, err := helper.AuthUsingLDAP(username, password)
	if !ok {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Error(w, "invalid username/password", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(data)
}
