package main

import "fmt"

func main() {
	anArray := [...]string{"Fadlan", "Abduh", "Erman", "Nawa"}

	/* ACCESSING ARRAY DATA TYPES
	i : index
	e : element
	*/

	// With for loop:
	fmt.Println("Accessing array data types using for loop")
	fmt.Println("i\\t: anArray[i]\n===============")

	for i := 0; i < len(anArray); i++ {
		fmt.Println(i, "\t:", anArray[i])
	}
	fmt.Println()

	// With for range:
	fmt.Println("Accessing array data types using for range")
	fmt.Println("i\\t: e\n===============")

	for i, e := range anArray {
		fmt.Println(i, "\t:", e)
	}
	fmt.Println()

	// For range without index
	// dengan menggunakan reserved variable untuk membuang index
	fmt.Println("For range without index")
	fmt.Println("e\n===============")

	for _, e := range anArray {
		fmt.Println(e)
	}
}
