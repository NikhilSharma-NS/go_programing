package main

import "fmt"

func main() {

	var fruits [5]string

	fruits[0] = "A"
	fruits[1] = "B"
	fruits[2] = "C"
	fruits[3] = "D"
	fruits[4] = "E"
	fmt.Println("Value Semantic")
	for i, value := range fruits {
		fmt.Println(i, value)
	}
	fmt.Println("Pointer Semantic")
	for i := range fruits {
		fmt.Println(i, fruits[i])
	}
	fmt.Println("Value Semantic Updation")
	//this will not update
	for i, value := range fruits {
		fruits[4] = "Z"
		fmt.Println(i, value)
	}
	fmt.Println("Pointer Semantic : updation")
	for i := range fruits {
		fruits[4] = "Z"
		fmt.Println(i, fruits[i])
	}
}
