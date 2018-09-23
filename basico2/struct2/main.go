package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	usr := user{"Vinicius", "vinicius@teste.com"}
	fmt.Println(usr)

	usrp := &usr
	usrp.name = "Vinicius Pacheco"
	fmt.Println(usrp)
}
