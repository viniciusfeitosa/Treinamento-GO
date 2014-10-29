package main

import "fmt"

const (
	x = 1
	y = 10
	z = "Olá"
)

const v string = "Hello World"

func main() {

	const u string = "Olá mundo" // Não usual

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(v)
	fmt.Println(u)
}
