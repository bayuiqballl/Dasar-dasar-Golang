package main

import (
	"fmt"
	"strconv"
)

func main() {
	//konstanta
	const gaji = 50
	//  konverter int to string
	gajiStr := strconv.Itoa(gaji)
	fmt.Println("hasilnya " + gajiStr)

	// konverter string to int
	gajiString := "10000"
	// undercover buat variabel tidak terpakai
	gaji2, _ := strconv.Atoi(gajiString)
	bonus := gaji2 + 5000

	fmt.Println("hasilnya ", bonus)

}
