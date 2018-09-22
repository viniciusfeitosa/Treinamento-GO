package main

import (
	"fmt"
)

func main() {
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init:", b)
	fmt.Println("b len:", len(b))
	fmt.Println("b cap:", cap(b))

	c := [...]int{1, 2, 3}
	fmt.Println("init:", c)
	fmt.Println("c len:", len(c))
	fmt.Println("c cap:", cap(c))
}
