package main

import (
	"fmt"
	"strconv"
)

func getBio(age int, name string, status string) string {
	ageNow := strconv.Itoa(age)
	return name + " adalah seorang " + status + " saat ini berumur " + ageNow + " tahun"

}

func main() {
	fmt.Println(getBio(21, "Bayu Muhammad Iqbal", "Bootcamp"))
	fmt.Println(getBio(88, "Ibrahim", "Nabi"))

}
