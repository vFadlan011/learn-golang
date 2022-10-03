package main

import "fmt"

func main() {
	abc := "abcdefghijklmnopqrstuvwxyz"
	// counter:= 0; counter <10   adalah init statement,
	// counter++ adalah post statement karena hanya dijalankan setiap loop berakhir
	for counter := 0; counter <= 26; counter++ {
		temp := ""

		// add alphabet[counter] to temp
		var abcCountera byte = abc[counter%26]
		abcCounter := string(abcCountera)
		temp += abcCounter

		// add alphabet[counter*2] to temp
		abcCountera = abc[(counter*2)%26]
		abcCounter = string(abcCountera)
		temp += abcCounter

		temp += "-"

		// add counter to temp
		strCounter := fmt.Sprint(counter)
		temp += strCounter

		// add counter*2 to temp
		strCounter = fmt.Sprint(counter * 2)
		temp += strCounter

		fmt.Println(temp)
	}
}
