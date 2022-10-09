package main

import "fmt"

func getHello(name string) (hello string) {
	hello += "Hello " + name + "!"
	return
}

func main() {
	sayHello := getHello
	fmt.Println(sayHello("fallan"))
}
