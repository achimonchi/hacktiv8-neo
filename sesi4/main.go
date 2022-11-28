package main

import (
	"sesi4/server"
	"sesi4/server/controllers"
)

func main() {
	PORT := ":9000"

	auth := controllers.NewAuthController()

	server.NewRouter(PORT, auth).Start()
}
