package main

import (
	"fmt"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're blocked")
	} else {
		fmt.Println("Welcome")
	}
}

func main() {
	// Storing a function into variable == Creating anonymous function
	blacklist := func(name string) bool {
		list := [...]string{"Adi", "Lukman", "Aldo", "xyz"}
		isBlocked := false

		for _, e := range list {
			if name == e {
				isBlocked = true
			}
		}
		return isBlocked
	}

	registerUser("Aldo", blacklist)
	registerUser("Eko Juniarto Khannedy", blacklist)
	/*
		anon function don't have to stored into var
		it also can be used as parameter directly (javascript: callback)
	*/
	registerUser("Lukman", func(name string) bool {
		list := [...]string{"Adi", "Lukman", "Aldo", "xyz"}
		isBlocked := false

		for _, e := range list {
			if name == e {
				isBlocked = true
			}
		}
		return isBlocked
	})

	/*
		IIFE == Imediately Invoked Function Expression
		Function that's immediately called after being declared
	*/
	fmt.Println("\n\n===IIFE===")
	numbers := []int{2, 3, 4, 6, 4, 8, 8, 2, 1, 6, 7, 7, 9, 5, 10, 8, 10, 6, 7, 8, 8, 10, 8, 8, 10}

	filteredNumbers := func(min int) []int {
		var filtered []int

		/*
			CLOSURE
			what closure do is calling data directly from other scope
			in this case it calls numbers from main scope
		*/
		for _, e := range numbers {
			if e < min {
				continue
			}
			filtered = append(filtered, e)
		}
		return filtered
	}(5)
	fmt.Println("numbers\t:", numbers)
	fmt.Println("filteredNumbers\t:", filteredNumbers)
}
