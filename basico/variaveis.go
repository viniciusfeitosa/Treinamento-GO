package main

import "fmt"

func main() {
	// Tipagem explicita
	var x string = "hello world"
	fmt.Println(x)

	// Tipagem explicita com atribuição posterior
	var y string
	y = "Olá Mundo"
	fmt.Println(y)

	// Tipagem implicita
	v := "Holla "
	z := "que tal"
	v += z // concatenando strings
	fmt.Println(v)

}
