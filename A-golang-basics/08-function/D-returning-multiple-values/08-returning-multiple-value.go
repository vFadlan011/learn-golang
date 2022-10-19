package main

import "fmt"

func powerOf(x int, n int) (string, int) {
	pow := x
	for i := 1; i < n; i++ {
		pow *= x
	}
	text := fmt.Sprint(x) + " power by " + fmt.Sprint(n) + " is"
	return text, pow
}

func printPow(x int, n int) {
	text, pow := powerOf(x, n)
	fmt.Println(text, pow)
}

func main() {
	printPow(2, 3)
	printPow(2, 4)
	printPow(4, 4)
	printPow(3, 4)
	printPow(9, 2)

	// ignore returned value using reserved var
	_, pow := powerOf(3, 3)
	fmt.Println(pow)
}
