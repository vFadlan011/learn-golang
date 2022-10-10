package main

import "fmt"

/*
	STRUCT CAN ACCOMODATE METHODS, just like a class
*/
type User struct {
	Name, Address string
	Age           int
}

func (user User) sayHi() {
	fmt.Println("Halo", user.Name+"👋")
}

func (user User) introduce(name string) {
	user.sayHi()
	fmt.Println("Perkenalkan, saya", name+"🤝")
}

func main() {
	eko := User{
		Name:    "Eko Kurniawan",
		Address: "Indonesia",
		Age:     30,
	}
	eko.introduce("Fadlan")
}
