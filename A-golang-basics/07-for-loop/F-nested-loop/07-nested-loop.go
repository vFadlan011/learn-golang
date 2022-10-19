package main

import "fmt"

func main() {

	// kita dapat melabeli loop dengan suatu nama,
	// kita dapat melakukan break, continue, dsb ke nama loop tersebut

loopLuar:
	for i := 1; i <= 10; i++ {
		fmt.Println()

		if i%2 == 0 {
			continue loopLuar
		}

		for j := 0; j <= 5; j++ {
			fmt.Print(i, ":", j, "  ")
			if j == 7 {
				break
			}
		}

	}
}
