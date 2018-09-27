package main

import (
	"errors"
	"fmt"
)

func falha(i int) (int, error) {
	if i < 10 {
		return -1, errors.New("O número tem que ser maior que 10")
	}
	return i, nil
}

func main() {
	if value, err := falha(11); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(value)
	}
}
