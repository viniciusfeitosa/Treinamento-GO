package main

import "fmt"

func algo() (string, int) {
	return "oito", 8
}

func main() {
	a, b := algo()
	fmt.Println(a, b)
}
