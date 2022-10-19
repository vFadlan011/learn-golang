package main

import "fmt"

func main() {
	/* konstanta digunakan untuk menyimpan variable yang nilainya
	tidak akan diubah di kemudian waktu. Sebagai contoh pi memiliki nilai 3.14,
	cara paling cepat untuk menyimpan nilai pi tersebut adalah dengan mendeklarasikan pi sebagai
	konstanta, karena nilai pi tidak akan pernah berubah.*/
	const pi = 3.14
	fmt.Println(pi)
	/* Deklarasi konstanta bisa dilakukan sekaligus*/
	const (
		square   = "kotak"
		box              //nilai box akan mengikuti variable sebelumnya, yaitu square yang bernilai "kotak"
		isCamel  bool    = true
		floatNum float32 = 2.56
	)
	fmt.Printf("square\t: \"%s\"\n", square)
	fmt.Printf("box\t: \"%s\"\n", box)
	fmt.Println("isCamel\t:", isCamel)
	fmt.Println("floatNum:", floatNum)
}
