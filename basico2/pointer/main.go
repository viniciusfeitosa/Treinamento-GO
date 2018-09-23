package main

import "fmt"

func pareceQueSiPero(val string) {
	val = "si"
}

func pareceQueNoPero(val *string) {
	*val = "si"
}

func main() {
	i := "no"
	fmt.Println("start:", i)

	pareceQueSiPero(i)
	fmt.Println("pareceQueSiPero:", i)

	pareceQueNoPero(&i)
	fmt.Println("pareceQueNoPero:", i)

	fmt.Println("pointer:", &i)
}
