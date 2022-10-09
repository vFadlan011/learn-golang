package main

import "fmt"

func sum(name string, item ...int) (text string, sum int) {
	text = "Total " + name + ":"

	for _, e := range item {
		sum += e
	}
	return
}

func main() {
	fmt.Println(sum("buah mangga", 2, 2, 2, 2))
	fmt.Println(sum("bola basket", 3, 1, 2, 4))

	sliceBata := []int{120, 96, 176, 224}
	fmt.Println(sum("bata hebel", sliceBata...))
}
