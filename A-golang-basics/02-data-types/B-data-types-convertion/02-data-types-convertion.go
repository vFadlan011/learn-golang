package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var des32 float32 = float32(nilai64) + 0.023

	fmt.Println("Nilai int32\t:", nilai32)
	fmt.Println("Nilai int64\t:", nilai64)
	fmt.Println("Nilai float32\t:", des32)

	var name = "Fadlan"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println("\nFadlan[0]\t:", e, "\t//ASCII Code of 'F'")
	fmt.Println("Fadlan[0]\t:", eString, "\t//e converted to string")
}
