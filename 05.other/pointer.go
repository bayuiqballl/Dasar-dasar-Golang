package main
import "fmt"

func changePoint(point *int){
	*point = 200
}


func main(){
	// Pointer
	var point = 100
	changePoint(&point)
	fmt.Println(point)
	fmt.Println("memori :",&point)
}