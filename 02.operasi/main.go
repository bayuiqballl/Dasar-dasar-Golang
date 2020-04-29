package main

import "fmt"

func main() {
	var hasil int

	// penjumlahan
	a := 10
	b := 20
	hasil = a * b
	fmt.Println("hasilnya ", hasil)

	// modulo
	var hasilModulo int
	aM := 32
	bM := 10
	hasilModulo = aM % bM
	fmt.Println("hasil Mod ", hasilModulo)
	// increment & decrement
	i := 1
	i++
	j := 2
	j--
	fmt.Println("hasil increment ", i, "hasil decrement", j)

}
