package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 9
	m["k2"] = 11

	fmt.Println("map:", m)

	v1 := m["k2"]
	fmt.Println("v1: ", v1)

	delete(m, "k2")
	fmt.Println("map:", m)
}
