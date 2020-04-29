package main

import "fmt"

func main() {
	// 1 + .. n (sekian)
	// 1 + 2 + 3 + 4 ... + 10 = 55

	total := 0
	for num := 1; num <= 10; num++ {
		total = total + num
	}

	fmt.Println("Hasil jumlah deret angka sampai n", total)

}
