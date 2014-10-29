package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Println(i)

		if i >= 10 {
			break
		}
		i += 1

	}

}
