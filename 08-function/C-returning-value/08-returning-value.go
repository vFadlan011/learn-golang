package main

import "fmt"

func substract(a int, b int) int {
	return a - b
}

func divide(a int, b int) int {
	if b == 0 {
		fmt.Print("Cannot divide by 0")
		return 0
	}

	return a / b
}

func main() {
	fmt.Println("Substract")
	fmt.Println(substract(12, 6))
	fmt.Println(substract(12, 12))
	fmt.Println(substract(6, 12))

	fmt.Println("\nDivide")
	fmt.Println(divide(3, 2))
	fmt.Println(divide(12, 4))
	fmt.Println(divide(12, 0))
}
