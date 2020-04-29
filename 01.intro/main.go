package main

import "fmt"

// Variable - Tipe Data String
var program string
var Pelajaran = "Go"

func main() {
	umur := 75

	// integer
	var total int
	a := 10
	b := 21
	total = a + b

	// short hand
	namaDepan, namaBelakang := "Bayu", "Iqbal"
	nama := namaDepan + " " + namaBelakang

	fmt.Println("Hello World " + Pelajaran)
	fmt.Println(nama+"umur saya", umur)
	fmt.Println(total)
}
