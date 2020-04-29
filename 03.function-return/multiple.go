// multiple return

package main

import (
	"fmt"
	"strconv"
)

func getBio(age int, name string, status string) (string, string) {
	ageNow := strconv.Itoa(age)
	return name + " adalah seorang " + status, " saat ini berumur " + ageNow + " tahun"

}

func main() {
	basicInfo, ageInfo := getBio(21, "Bayu Muhammad Iqbal", "Software Enginer")
	fmt.Println(basicInfo, ageInfo)
}
