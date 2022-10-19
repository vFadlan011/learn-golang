package main

import "fmt"

func getName() (firstName string, middleName string, lastName string, alterName string) {
	firstName = "Fadlan"
	middleName = "Abduh"
	lastName = "Erman"
	alterName = "Fallan"
	// cukup menuliskan return, jika tidak akan mengubah urutan return valuenya
	// Urutan variable yang di-return dapat diacak
	return middleName, lastName, alterName, firstName
}

func main() {
	fmt.Println(getName())
}
