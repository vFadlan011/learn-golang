package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func printFactorial(value int, factor func(int) int) {
	fmt.Println(factor(value))
}

func main() {
	printFactorial(2, factorialLoop)
	printFactorial(3, factorialLoop)
	printFactorial(4, factorialLoop)
	printFactorial(5, factorialLoop)
	fmt.Println()
	printFactorial(2, factorialRecursive)
	printFactorial(3, factorialRecursive)
	printFactorial(4, factorialRecursive)
	printFactorial(5, factorialRecursive)
}
