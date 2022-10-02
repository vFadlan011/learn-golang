package main

import "fmt"

func main() {
	var arr [4]string
	arr[0] = "Eko"
	arr[1] = "Kurniawan"
	arr[2] = "Khannedy"

	fmt.Println("arr\t:", arr)
	fmt.Println("arr[0]\t:", arr[0])
	fmt.Println("arr[1]\t:", arr[1])
	fmt.Println("arr[2]\t:", arr[2])
	fmt.Println("arr[3]\t:", arr[3])
	fmt.Println("len(arr):", len(arr))
	fmt.Println()

	var arr2 = [5]int{
		90,
		86,
		93,
		96,
		84,
	}

	fmt.Println(arr2)
	fmt.Println("len(arr2):", len(arr2))
}
