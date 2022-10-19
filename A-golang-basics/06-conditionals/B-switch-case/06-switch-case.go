package main

import "fmt"

func main() {
	/* Basic switch case */
	fmt.Println("#Basic Switch Case")
	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hello Brow\nKenalan dunk")
	}

	/* Short Statement switch case */
	fmt.Println("\n#Switch Case Short Statement")
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	/* SWITCH CASE WITH IF-ELSE STYLE*/
	fmt.Println("\n#Switch Case with if-else Style")
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
	fmt.Println("\n#falltrough Keyword")
	name = "Eko Kurniawan"
	switch {
	case len(name) > 7:
		fmt.Println("Nama tidak valid")
		fallthrough // falltrough hanya melanjut 1 case selanjutnya
	case len(name) >= 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
