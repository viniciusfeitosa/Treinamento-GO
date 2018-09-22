package main

import "fmt"

func main() {
	s := []string{"v", "i", "n"}

	s = append(s, "i")
	s = append(s, "c", "i", "u", "s")

	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)
}
