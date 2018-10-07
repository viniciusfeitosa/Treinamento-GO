package main

import "fmt"

func main() {
	msg := make(chan string, 2)

	msg <- "ping"
	msg <- "pong"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}
