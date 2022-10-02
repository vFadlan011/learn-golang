package main

import "fmt"

func main() {
	/* SWITCH CASE WITH CONDITION */
	var score int8 = 90
	var point, msg string

	switch {
	case (score <= 100) && (score > 95):
		point = "A+"
		msg = "Perfect!"

	case (score <= 95) && (score > 89):
		point = "A"
		msg = "Awesome!"

	case (score <= 89) && (score > 81):
		point = "B"
		msg = "Keep it! And go higher!"

	case (score <= 81) && (score > 75):
		point = "C"
		msg = "Not bad, you need to learn more"

	case (score == 75):
		point = "D"
		msg = "You need to learn more"

	case (score < 75):
		point = "E"
		msg = "Learn more and more!"

	default:
		point = "NULL!"
	}

	fmt.Println(score, point)
	fmt.Println(msg)

	/* SWITCH CASE & falltrough KEYWORD	 */

}
