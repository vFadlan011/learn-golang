package main

import "fmt"

func addition(a int, b int) {
	fmt.Println(a, "+", b, "\t:", a+b)
}

func main() {
	addition(1, 1)
	addition(1, 2)
	addition(1, 3)
	addition(1, 4)
}
