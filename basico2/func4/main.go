package main

import "fmt"

func add(inputs ...int) (total int) {
	for _, v := range inputs {
		total += v
	}
	return
}

func main() {
	a := add(1, 2, 3)
	fmt.Println(a)
	b := add([]int{1, 2, 3}...)
	fmt.Println(b)
}
