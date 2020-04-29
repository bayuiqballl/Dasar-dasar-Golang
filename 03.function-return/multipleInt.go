package main

import "fmt"


func getBio(age int, name string, status string) (bio string, ageIn10 int) {
	ageIn10 = age + 10
	bio = name + " adalah seorang " + status
	return
}

func main() {
	basicInfo, ageInfo := getBio(21,"Bayu Muhammad Iqbal", "Software Engineer")

	fmt.Println(basicInfo)
	fmt.Println("umurnya 10 tahun lagi adalah", ageInfo)

}