package main

import "fmt"

/*
	Struct is just like a map that can accomodate different data types
	(python: dictionary, oop: object and properties)
*/
type User struct {
	Name, Address string
	Age           int
}

func main() {
	var fadlan User
	fadlan.Name = "Fadlan Abduh"
	fadlan.Address = "Jawa Tengah"
	fadlan.Age = 16

	fmt.Println(fadlan)

	// STRUCT LITERALS
	eko := User{
		Name:    "Eko Kurniawan Khannedy",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(eko)

	/*
		STRUCT LITERALS SHORTHAND
		joko := User{"Joko Kurniawan", 30, "Indonesia"}
		will cause an error because of wrong declaration order

		User : Name, Address, Age
		joko : Name, Age, Address

		THIS SHORTHAND ISN'T RECOMENDED
		can cause an error if declared with wrong order
	*/
}
