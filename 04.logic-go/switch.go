package main

import "fmt"

func bayarUtang(uang int) (pesan string) {
	utang := 50000

	switch {
	case uang > utang:
		pesan = "Uang kamu kebanyakan"
	case uang == utang:
		pesan = "Uangya pas, lunas!"
	case uang < utang:
		pesan = "Uangnya tidak cukup"
	}

	return
}

func main() {
	fmt.Println(bayarUtang(560000))
}
