package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1.0
	b := 2.0
	c := a + b          //3
	d := a + c          //4
	e := 10 / 2         //5
	f := b * c          //6
	g := 14 - (14 / 2)  //7
	h := math.Pow(b, c) //8

	fmt.Println("1 + 3\t\t:", d)
	fmt.Println("10 / 2\t\t:", e)
	fmt.Println("2 * 3\t\t:", f)
	fmt.Println("14 - (14/2)\t:", g)
	fmt.Println("math.Pow(b,c)\t:", h)
}
