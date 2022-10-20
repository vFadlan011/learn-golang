package main

import "fmt"

func main() {
	// Anonymous struct
	// We can create a struct that just used once

	// Struct initialization without defining it's own properties directly
	dimas := struct {
		name   string
		gender string
		age    int
		grade  int
	}{}

	dimas.name = "Dimas Wijayanto"
	dimas.gender = "M"
	dimas.age = 17
	dimas.grade = 2

	// Struct initialization with properties defined directly
	lucy := struct {
		name   string
		gender string
		age    int
		grade  int
	}{
		name:   "Lucy Despita",
		gender: "F",
		age:    17,
		grade:  2,
	}

	fmt.Println(dimas)
	fmt.Println(lucy)
}
