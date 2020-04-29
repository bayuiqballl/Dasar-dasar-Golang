package main

import "fmt"

//  fungsi dan parameter

func multiply(angka1 int, angka2 int) int {
	return angka1 * angka2
}

func main() {
	fmt.Println("Fungsi ini menghasilkan", multiply(10, 5))
	fmt.Println("Fungsi yang sama  ini menghasilkan", multiply(4, 5))
}
