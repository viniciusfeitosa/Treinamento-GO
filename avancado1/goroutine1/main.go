package main

import (
	"fmt"
	"time"
)

func loop(id, input int) {
	for i := 0; i <= input; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("From id:", id, "the value is:", i)
	}
}

func main() {
	go loop(1, 10)
	loop(2, 10)
	fmt.Scanln()
	fmt.Println("Done")
}
