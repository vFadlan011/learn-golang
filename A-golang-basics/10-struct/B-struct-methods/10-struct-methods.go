package main

import "fmt"

type Cat struct {
	Name  string
	Color string
	Age   int
}

// To create a method, we just create a function with receiver argument
// Receiver Argument
//    \/  \/
func (c Cat) meow() {
	fmt.Println(c.Name, "\t: Meoww")
}

type Duck struct {
	Name   string
	Age    int
	Weight int
}

func (d Duck) quack() {
	fmt.Println(d.Name, "\t: Quack")
}

func main() {
	dudul := Cat{
		Name:  "Dudul",
		Color: "Orange",
		Age:   6,
	}
	fmt.Println(dudul.Name)
	dudul.meow()
	kwek := Duck{
		Name:   "Kwek",
		Age:    7,
		Weight: 4,
	}
	kwek.quack()
}
