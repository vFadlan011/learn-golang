package main

import "fmt"

func main() {
	score := 82
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
	} else if score < 75 { //0-74
		point = "E"
	}
	fmt.Print(score, point)
}
