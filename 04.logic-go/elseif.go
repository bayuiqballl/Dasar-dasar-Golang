package main

import "fmt"

func main() {
	utang := 10000
	uang := 10000

	if uang > utang {
		fmt.Println("Utang lunas bosku")
	} else if uang == utang {
		fmt.Println("uangnya pas, lunas")
	} else {
		fmt.Println("Uang kamu tidak cukup")
	}
}
