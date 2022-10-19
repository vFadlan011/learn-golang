package main

import "fmt"

func main() {
	type id string
	type married bool

	var idSaya id = "rj-2368"
	var ismarried married = false

	fmt.Println(idSaya, ismarried)
}
