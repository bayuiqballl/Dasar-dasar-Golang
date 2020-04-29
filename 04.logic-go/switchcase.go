package main

import "fmt"

func cekLampu(warna string) (pesan string) {

	switch warna {
	case "merah":
		pesan = "Berhenti!"
	case "kuning":
		pesan = "Hati-hati!"
	case "hijau":
		pesan = "Jalan!"
	default:
		pesan = "Lampu sedang error.."
	}

	return
}

func main() {
	fmt.Println(cekLampu("hijau"))
}
