package main

import "fmt"

func main() {
	var score int = 100

	var point string
	if score <= 100 && score > 95 { //96-100
		point = "A+"
	} else if score <= 95 && score > 89 { //90-95
		point = "A"
	} else if score <= 89 && score > 81 { //82-89
		point = "B"
	} else if score <= 81 && score > 75 { //76-81
		point = "C"
	} else if score == 75 { //75
		point = "D"
	} else if score < 75 && score > 39 { //0-74
		point = "E"
	} else if score <= 39 && score > 10 {
		point = "F"
	} else {
		point = "ABSOLUTE-IDIOTS!"
	}
	fmt.Println("Basic")
	fmt.Println(score, point)

	name := "Fadlan"

	fmt.Println("\nShort Statement")
	if length := len(name); length <= 5 {
		fmt.Println("Format Nama Benar")
	} else {
		fmt.Println("Nama terlalu panjang")
	}

}
