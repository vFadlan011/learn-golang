package main

import "fmt"

type HasName interface {
	GetName() string
}

type User struct {
	Name string
}
type Animal struct {
	Name string
}

func (user User) GetName() string {
	return user.Name
}
func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	eko := User{"Eko Kurniawan"}
	sayHello(eko)

	kusing := Animal{"Kusing"}
	sayHello(kusing)
}
