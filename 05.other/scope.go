package main

import "fmt"

var bonus = 50

// scope variable
func hitungTotal(gaji int) int {
	return gaji + bonus
}

func main() {
	fmt.Println(hitungTotal(100))

	bonus = 100
	fmt.Println("Bonus =", bonus)
}
