package server

import (
	"log"
	"net/http"
	"sesi4/server/controllers"
)

type Router struct {
	port string
	auth *controllers.AuthController
}

func NewRouter(port string, auth *controllers.AuthController) *Router {
	return &Router{
		port: port,
		auth: auth,
	}
}

func (r *Router) Start() {

	http.HandleFunc("/", r.auth.HandleIndex)
	http.HandleFunc("/login", r.auth.HandleLogin)

	log.Println("Server running at port", r.port)
	http.ListenAndServe(r.port, nil)
}
