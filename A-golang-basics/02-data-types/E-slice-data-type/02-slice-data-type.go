package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"Augustus",
		"September",
		"October",
		"November",
		"December",
	}

	aprAug := months[3:8]
	fmt.Println("Months[3:8]\t:", aprAug)
	fmt.Println("len(aprAug)\t:", len(aprAug), "\ttotal ada 5 bulan")
	fmt.Println("cap(aprAug)\t:", cap(aprAug), "\tApril sampai December adalah 9 bulan")
	fmt.Println()

	/*
		slice mengacu kepada array, jadi jika array berubah maka slice pun ikut berubah
		months[5] = "Diubah"
		fmt.Println(aprAug)

		begitu pula sebaliknya, jika merubah slice maka array juga ikut berubah
		aprAug[0] = "April lagi"
		fmt.Println(months) //January February March April lagi ...
	*/

	/* METHOD PADA SLICE */

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	/* APPEND: Menambahkan data pada slice ke id terakhir, tanpa menambahkannya ke array utama
	Meskipun slice sudah memenuhi capacity, jika kita menggunakan append maka akan dibuatkan array baru
	yang direferensikan oleh slice.
	Jadi meskipun kita merubah data pada slice tersebut, maka array utama tidak akan ikut berubah. */

	fmt.Println("\nAppend slice yang membuat array baru")
	daysSlice1 := append(days[5:], "Libur-baru")
	fmt.Println(daysSlice1)
	daysSlice1 = append(daysSlice1, "Libur-baru-lagi")
	fmt.Println(daysSlice1)
	daysSlice1[0] = "Bukan-Sabtu"
	fmt.Println(daysSlice1)
	fmt.Println(days)
	fmt.Println()

	/*
		###PENTING###

		Jika slice yang kita append masih sanggup menampung data, maka slice tersebut tidak akan
		membuatkan array baru. Jadi jika kita mengubah nilai pada data tersebut, maka data yang ada pada
		array utama akan ikut berubah.
	*/
	fmt.Println("\nAppend slice yang tidak membuat array baru")
	daysSlice2 := append(days[:2], "Libur-tengah-minggu")
	fmt.Println(daysSlice2)
	fmt.Println(days)
	fmt.Println()

	/* MAKE SLICE */
	// make([]type, len, cap)
	newSlice := make([]string, 3) // len:3; cap:3
	newSlice[0] = "Red"
	newSlice[1] = "Green"
	newSlice[2] = "Blue"

	fmt.Println("\nCreate Slice with 'make'")
	fmt.Println("newSlice\t:", newSlice)
	fmt.Println("len(newslice)\t:", len(newSlice))
	fmt.Println("cap(newSlice)\t:", cap(newSlice))
	fmt.Println()

	/* COPY SLICE */
	// copy(destination, origin)
	copySlice := make([]string, 3, 5) // len:3; cap:2
	copy(copySlice, newSlice)

	fmt.Println("\nCopy Slice with 'copy'")
	fmt.Println("copySlice\t:", copySlice)
	fmt.Println("len(copySlice)\t:", len(copySlice))
	fmt.Println("cap(copySlice)\t:", cap(copySlice))
	fmt.Println()

	fmt.Println("Appending 'Purple' and 'Yellow' to copySlice")
	copySlice = append(copySlice, "Purple")
	copySlice = append(copySlice, "Yellow")
	fmt.Println("copySlice\t:", copySlice)
	fmt.Println()

	/* Membuat slice tanpa menggunakan keyword 'make' */
	iniSlice := []int{32, 33, 34, 35}
	fmt.Println("\nCreate slice without keyword 'make'")
	fmt.Println("iniSlice\t:", iniSlice)
	fmt.Println("len(iniSlice)\t:", len(iniSlice))
	fmt.Println("cap(iniSlice)\t:", cap(iniSlice))
	fmt.Println()
}
