package main

import "fmt"

func endApp() {
	fmt.Println("Program ended")
}

func runApp(error bool) {
	defer endApp()
	if error {
		/*
			when panic function called
			=> defer function still called
				 code below panic, will be skipped (in for loop: break)
		*/
		panic("AN ERROR OCCURED")
	}
	fmt.Println("Program runing...")
}

func main() {
	runApp(true)
}
