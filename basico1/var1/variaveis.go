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
	var n = "Olá Mundo"
	fmt.Println(n)

	// Tipagem implicita dentro de uma função
	v := "Holla "
	z := "que tal"
	v += z // concatenando strings
	fmt.Println(v)

}
