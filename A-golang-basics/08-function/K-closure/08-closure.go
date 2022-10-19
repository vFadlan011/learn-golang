package main

import "fmt"

func main() {
	/*
		CLOSURE
		what closure do is calling data directly from other scope
		in this case it calls counter from main scope
	*/
	name := "Eko"
	counter := 0
	increment := func() {
		// it isn't a closure, bcs it's declare a new variable
		name := "Fadlan"
		fmt.Println("Increment")
		fmt.Println(counter)
		fmt.Println(name)
		counter++
	}
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
