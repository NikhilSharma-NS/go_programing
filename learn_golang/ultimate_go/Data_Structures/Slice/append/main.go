package main

import "fmt"

func main() {

	slice := make([]string, 3, 3)
	slice[0] = "A"
	slice[1] = "B"
	slice[2] = "C"
	//Below will give Runtime error
	//slice[3] = "D"

	fmt.Println("Len:Cap", len(slice), cap(slice)) //3 3
	slice = append(slice, "D")

	fmt.Println("Len:Cap", len(slice), cap(slice)) //4 6

}
