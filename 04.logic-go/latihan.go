package main

import "fmt"

func bayarUtang(uang int) (pesan string) {
	utang := 50000
	if uang > utang {
		pesan = "Uang kamu kebanyakan"
	} else if uang == utang {
		pesan = "Lunass"
	} else {
		pesan = "Uang kamu tidak cukup"
	}

	return
}

func main() {
	fmt.Println(bayarUtang(50000))
}
