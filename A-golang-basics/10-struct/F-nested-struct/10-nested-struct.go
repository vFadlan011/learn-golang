package main

import "fmt"

type Book struct {
	title     string
	author    string
	publisher string
	year      int
}

type Novel struct {
	Book
	genre string
}

func main() {
	IRobot := Novel{
		Book: Book{
			title:     "I, Robot",
			author:    "Isaac Asimov",
			publisher: "",
			year:      1950,
		},
		genre: "sci-fi",
	}
	// it returns the same thing
	fmt.Println(IRobot.title)
	fmt.Println(IRobot.Book.title)
}
