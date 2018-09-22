package main

import "fmt"

func main() {
	a := []string{"v", "i", "n"}
	b := []string{"c", "i", "u", "s"}
	fmt.Println(a, b)

	a = append(a, b...)
	fmt.Println(a)

	b = append(a[:2], a[3:]...)
	fmt.Println(b)

	// Leia para mais operações https://github.com/golang/go/wiki/SliceTricks
}
