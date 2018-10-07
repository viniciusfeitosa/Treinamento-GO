package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <- fmt.Sprintf("%s, pong", <-pings)
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "ping")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
