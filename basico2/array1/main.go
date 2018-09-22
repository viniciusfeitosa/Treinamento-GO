package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("start:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("limit:", a[3:5])

	fmt.Println("len:", len(a))
}
