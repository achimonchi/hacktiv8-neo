package main

import "sesi5/server"

const APP_Port = ":5555"

func main() {
	server.NewRouter(APP_Port).Start()
}
