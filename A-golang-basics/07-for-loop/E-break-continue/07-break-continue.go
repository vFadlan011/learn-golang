package main

import "fmt"

func main() {
	/*
		- Keyword continue akan memaksa for loop ke iterasi selanjutnya
		- Keyword break akan memberhentikan for loop secara paksa
	*/
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
			// perintah selanjutnya tidak akan dijalankan
			fmt.Println("genap dilewati")
		}
		if i >= 8 {
			fmt.Println("loop berakhirr")
			break
		}
		fmt.Println(i)
	}
}
