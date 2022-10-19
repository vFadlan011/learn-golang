package main

import "fmt"

type ArrINT []int
type str string

// Methods can also declared for non struct type
func (a ArrINT) GetAVG() float64 {
	sum := 0
	for _, e := range a {
		sum += e
	}
	avg := float64(sum) / float64(len(a))
	return avg
}

func (s str) lower() string {
	temp := ""
	for _, e := range s {
		if e <= 90 && e >= 65 {
			temp += string(e + 26 + 6)
		} else {
			temp += string(e)
		}
	}
	return temp
}

func main() {
	a := ArrINT{1, 9, 2, 8, 3, 3, 6, 3, 1, 1, 8, 7, 4, 6, 5}
	fmt.Println(a.GetAVG())

	var b str = "Lorem Ipsum Dolor Sit Amet"
	fmt.Println(b.lower())
}
