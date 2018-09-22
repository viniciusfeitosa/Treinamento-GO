package main

import "fmt"

func add(inputs ...int) {
	var total int
	for _, v := range inputs {
		total += v
	}
	fmt.Println(total)
}

func main() {
	add(1, 2, 3)
	add([]int{1, 2, 3}...)
}
