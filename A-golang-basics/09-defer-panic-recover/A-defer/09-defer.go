package main

import "fmt"

func endApp() {
	fmt.Println("Program ended.")
}

func app(value int) {
	/*
		with defer, go will run endApp() at the end of this function,
		even if there is an error
	*/
	defer endApp()
	result := 10 / value
	fmt.Println("Program running...")
	fmt.Println(result)
}

func main() {
	app(0)
}
