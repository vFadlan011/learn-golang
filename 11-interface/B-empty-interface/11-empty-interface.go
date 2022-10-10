package main

import "fmt"

type any interface {
}

func ups(a any) interface{} {
	return "Ups"
}

func main() {
	fmt.Println(ups(20))
}
