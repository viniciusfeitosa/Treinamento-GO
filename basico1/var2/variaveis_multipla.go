package main

import (
	"fmt"
	"reflect"
)

var (
	x         = "Hello "
	y         = "world"
	tipoInt   = 1
	tipoFloat = 2.0
)

func main() {
	fmt.Println(x, y)
	fmt.Printf("Tipo: %s, valor: %d \n", reflect.TypeOf(tipoInt).Kind(), tipoInt)
	fmt.Printf("Tipo: %s, valor: %.2f", reflect.TypeOf(tipoFloat).Kind(), tipoFloat)
}
