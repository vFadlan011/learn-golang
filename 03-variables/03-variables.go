package main

import "fmt"

func main() {
	// deklarasi variable dengan menuliskan tipe data secara eksplisit
	var firstName string = "john"
	// deklarasi variable tanpa menuliskan tipe data
	var lastName = "doe"
	// deklarasi variable tanpa menggunakan var
	fullName := firstName + " " + lastName
	fmt.Println(firstName, lastName, fullName)

	// operasi assignment
	lastName = "dane"
	fmt.Println(lastName)

	/* Deklasrasi multi-variable
	dapat dilakukan dengan menuliskan atau tidak menuliskan tipe data secara eksplisit
	tidak perlu menuliskan tipe data secara eksplisit jika dalam satu baris deklarasi
	terdapat tipe data yang berbeda */
	var the, quick, fox, two, threePointThree = "the", "quick", "fox", 2, 3.3
	fmt.Println(the, quick, fox, two, threePointThree)
}
