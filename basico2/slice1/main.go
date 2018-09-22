package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("content:", s)
	fmt.Println("cap:", cap(s))

	s[0] = "v"
	s[1] = "i"
	s[2] = "n"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
}
