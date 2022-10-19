package main

import "fmt"

func main() {
	// dalam go kita bisa membuat for loop tanpa menggunakan statement apapun
	// namun for loop seperti ini sulit dipahami, lebih baik gunakan statement saja
	counter := 0

	for {
		fmt.Println(counter)

		counter++
		if counter == 10 {
			break
		}
	}
}
