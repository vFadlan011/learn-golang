package main

import "fmt"

func main() {
	fmt.Println(1, "\t: satu\t\t\t[int]")
	fmt.Println(2, "\t: dua\t\t\t[int]")
	fmt.Println("###===###===###===###===###===###===###")
	fmt.Println(3.5, "\t: tiga koma lima\t[float]")
	fmt.Println("###===###===###===###===###===###===###")
	fmt.Println(true, "\t: true/benar\t\t[bool]")
	fmt.Println(false, "\t: false/salah\t\t[bool]")
	fmt.Println("###===###===###===###===###===###===###")
	fmt.Println("kalimat", ": \"kalimat\"\t\t[string]")
	fmt.Println("kalimat"[1], "\t: \"kalimat\"[1]\t\t[int]")
	fmt.Println(len("kalimat"), "\t: len(\"kalimat\")\t[int]")
}
