package main

import "fmt"

type struct1 struct {
	internalStruct struct2
}

type struct2 struct {
	value int
}

func main() {
	fmt.Println(
		struct1{
			struct2{value: 12},
		},
	)
}
