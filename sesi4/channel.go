package main

import (
	"fmt"
)

func main() {
	msg := make(chan string)
	msgErr := make(chan error)
	go goroutine(msg, msgErr)

	for {
		select {
		case err := <-msgErr:
			fmt.Println("ada error nih :", err)
			return
		case m := <-msg:
			fmt.Println("success", m)
			return
		}
	}
}

func goroutine(msg chan string, msgErr chan error) {
	fmt.Println("Ini dari goroutine")

	// msgErr <- fmt.Errorf("ini error dari goroutine")
	msg <- "Ini success dari goroutine"
}
