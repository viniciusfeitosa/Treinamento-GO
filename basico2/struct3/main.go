package main

import "fmt"

type user struct {
	name, email string
}

func (u user) changeEmail(newEmail string) string {
	u.email = newEmail
	return u.email
}

func (u *user) changeName(newName string) string {
	u.name = newName
	return u.name
}

func main() {
	usr := user{"Vinicius", "vinicius@teste.com"}
	fmt.Println(usr)
	fmt.Println(usr.changeEmail("bla@bla.com"))
	fmt.Println(usr)

	usrp := &usr
	fmt.Println(usrp)
	fmt.Println(usrp.changeName("Vinicius Pacheco"))
	fmt.Println(usrp)
}
