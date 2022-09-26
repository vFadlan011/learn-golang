package main

import "fmt"

func main() {
	// deklarasi variable dengan menuliskan tipe data secara eksplisit
	var firstName string = "john"
	// deklarasi variable tanpa menuliskan tipe data (disebut type inference, untuk best practicenya di google banyak)
	var lastName = "doe"
	// deklarasi variable tanpa menggunakan var
	fullName := firstName + " " + lastName
	fmt.Println(firstName, lastName, fullName)

	// operasi assignment
	lastName = "dane"
	fmt.Println(lastName)

	// Deklasrasi multi-variable
	var five, six, seven int8 = 5, 6, 7
	fmt.Println(five, six, seven)
	/*
		dapat dilakukan dengan menuliskan atau tidak menuliskan tipe data secara eksplisit
		tidak perlu menuliskan tipe data secara eksplisit jika dalam satu baris deklarasi
		terdapat tipe data yang berbeda */
	var the, quick, fox, two, threePointThree = "the", "quick", "fox", 2, 3.3
	fmt.Println(the, quick, fox, two, threePointThree)

	// Deklarasi multi-variable yang lebih ringkas
	today, is, dayName, date, monthName := "today", "is", "Tuesday,", 27, "sept"
	fmt.Println(today, is, dayName, date, monthName)

	// variable _ dalam go
	// digunakan untuk menampung nilai yang tidak terpakai
	_ = "something that useless"
}
