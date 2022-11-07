package main

import log "github.com/NooBeeID/go-logging/messages"

func main() {
	print("Ini dari main.go")
}

func print(text string) {
	log.SysInfo(text)
}
