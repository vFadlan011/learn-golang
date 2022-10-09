package main

import "fmt"

func filterMax(max int, item ...int) ([]int, func() int, func() []int, func(int), func(int)) {
	var res []int
	for _, e := range item {
		if e <= max {
			res = append(res, e)
		}
	}

	/*
		returned function
			getLen() and getData() return a closure.
			It means, every time we pushData() or rmAllData(),
			value that returned from getLen() or getData() will be changed

			IMO it's like an OOP, the parameter on main function is a constructor. Instead of calling method with object.method(), we use and treat returned function as a method
	*/

	getLen := func() int {
		return len(res)
	}
	getData := func() []int {
		return res
	}
	pushData := func(data int) {
		if data <= max {
			res = append(res, data)
			fmt.Println("\nPushing", data, "to array")
		} else {
			fmt.Println("\nUnable to push", data, "because it's higher than the limit", max)
		}
	}
	rmAllData := func(data int) {
		var resTemp []int
		var counter int

		fmt.Println()
		for _, e := range res {
			if e != data {
				resTemp = append(resTemp, e)
			} else {
				counter++
				fmt.Println("Deleted", counter, "items")
			}
		}

		res = resTemp
	}

	return res, getLen, getData, pushData, rmAllData
}

func main() {
	data, len, getData, pushData, rmAllData := filterMax(7, 2, 4, 3, 69, 585, 3, 4, 6, 7, 8, 8, 4, 5, 2, 7, 7, 4, 4, 3, 5, 8, 8, 6, 9, 9, 89)
	show := func() {
		fmt.Println("\nlen\t\t:", len())
		fmt.Println("getData\t\t:", getData())
		fmt.Println("data (no update):", data)
	}

	//
	show()
	pushData(4)
	show()
	rmAllData(2)
	show()
	pushData(7)
	pushData(9)
	pushData(3)
	show()
}
