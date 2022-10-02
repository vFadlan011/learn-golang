package main

import "fmt"

/*
Map pada go sama seperti object pada javascript atau dictionary pada python
Map adalah type data yang berisi key-value pair, key harus bersifat unik (hanya boleh ada 1 dalam sebuah map)
Bedanya, pada map golang type data pada key ataupun value tidak boleh berubah-ubah
*/
func main() {
	// Pembuatan map
	person := map[string]string{
		"name":"Fadlan Abduh",
		"address":"Pemalang",
	}

	
	// Menambahkan key dan value baru pada map
	person["title"] = "Programmer"
	fmt.Println(person)
	
	
	// Mendapatkan value diketahui key (indexing)
	fmt.Println("\nIndexing map")
	fmt.Println("person['name']\t:", person["name"]) // mengembalikan value dari key 'name'
	fmt.Println("person['abc']\t:", person["abc"], "\t//'abc' is not defined") // tidak mengembalikan apapun
	
	
	// Mendapatkan jumlah data pada map
	// len(map)
	fmt.Println("\nGet map length")
	fmt.Println("len(person)\t:",len(person))
	
	
	// Menghapus key-value pada map
	// delete(map, key)
	fmt.Println("\nDeleting key-value inside math using delete()")
	delete(person, "title")
	fmt.Println(person)


	// Membuat map dengan keyowrd 'make'
	// make(map[TypeKey]TypeValue)
	book := make(map[string]string)
	book["title"] = "Golang for Idiots"
	book["author"] = "John Doe"

	fmt.Println("\nCreate map with keyword 'make'")
	fmt.Println(book)
}