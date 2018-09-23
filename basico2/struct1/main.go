package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	usr := user{"Vinicius", "vinicius@teste.com"}
	fmt.Println(usr)

	usr = user{name: "Vinicius", email: "vinicius@teste.com"}
	fmt.Println(usr)

	usr = user{name: "Vinicius"}
	fmt.Println("name:", usr.name, "| email:", usr.email)

	usr = user{email: "vinicius@teste.com"}
	fmt.Println("name:", usr.name, "| email:", usr.email)
}
