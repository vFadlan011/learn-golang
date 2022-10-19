package main

import "fmt"

/*
	Struct is just like a class,
	but it just declare a property (in declaration block)
	and of course doesn't have a constructor method
*/
type User struct {
	Username string
	email string
	Followers int
	Following int	
}

func main() {
	fallan := User{}

	fallan.Username = "fallan"
	fallan.email = "fallnfall@gmail.com"
	fallan.Followers = 2048
	fallan.Following = 128

	eko := User{
		Username: "Eko Kurniawan",
		email: "Ekokhann@gm.co",
		Followers: 8192,
		Following: 128,
	}

	fmt.Println(fallan)
	fmt.Println(eko)
}