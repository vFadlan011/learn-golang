package main

import (
	"fmt"
	"public-private-properties/lib"
)

func main() {
	data1 := lib.MakeStats([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(data1)
	data1.Append(1)
	fmt.Println(data1)
}
