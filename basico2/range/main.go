package main

import "fmt"

func main() {

	for k, v := range "Vinicius" {
		fmt.Println(k, v, string(v))
	}
}
