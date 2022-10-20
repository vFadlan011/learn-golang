package main

import "fmt"

type Person struct {
	name   string
	gender string
	age    int
	job    string
}

func main() {
	// we can assigning an object into a list
	companyA := []Person{
		{name: "Khalid Khawarizmi", gender: "M", age: 30, job: "UI/UX"},
		{name: "Dwi Artanto", gender: "M", age: 27, job: "IT Operations"},
		{name: "Muhammad Sumbul", gender: "M", age: 24, job: "Junior Backend Engineer"},
	}

	companyB := []struct {
		name   string
		gender string
		age    int
		job    string
	}{
		{name: "Samsul Arip", gender: "M", age: 37, job: "Product Manager"},
		{name: "Lily Floyd", gender: "F", age: 28, job: "UX Researcher"},
		{name: "Andi Lukito", gender: "M", age: 42, job: "Principal Backend Engineer"},
	}

	PrintList := func() {
		for _, e := range companyA {
			fmt.Println(e)
		}
		fmt.Println("\n\n\n###############")
		for _, e := range companyB {
			fmt.Println(e)
		}
	}
	PrintList()
}
