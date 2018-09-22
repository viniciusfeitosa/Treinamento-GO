package main

import "fmt"

func main() {
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	for k, v := range n {
		fmt.Println(k, v)
	}

	for k := range n {
		fmt.Println("key:", k)
	}

	for _, v := range n {
		fmt.Println("Value:", v)
	}
}
