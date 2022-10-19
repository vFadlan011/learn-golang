package main

import "fmt"

func askInt(question string) (answer int) {
	fmt.Print(question)
	fmt.Scanln(&answer)
	return
}

func divide(nominator int, denominator int) {
	if denominator == 0 {
		panic("Cannot divide by 0")
	}

	fmt.Println(fmt.Sprint(nominator)+"/"+fmt.Sprint(denominator)+"\t\t:", nominator/denominator)
}

func endApp() {
	/*
		recover() work by catching an error message,
		if there wasn't an error => message = nil
	*/
	message := recover()
	if message != nil {
		fmt.Println("\nProgram ended because an error.")
		fmt.Println("An Error Occured:", message)
	} else {
		fmt.Println("Program finished.")
	}
}

func App() {
	defer endApp()
	nominator := askInt("Pembilang\t: ")
	denominator := askInt("Penyebut\t: ")
	divide(nominator, denominator)

}

func main() {
	App()
}
