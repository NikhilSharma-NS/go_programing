package main

import "fmt"

func main() {

	card_slice := []string{newCard()}
	fmt.Println("slice", card_slice)
	card_slice = append(card_slice, newCard())
	card_slice = append(card_slice, newCard())
	fmt.Println("slice_after_add", card_slice)

	card_array := [2]string{newCard()}
	fmt.Println("array", card_array)
	card_array[1] = newCard()
	//	card_array[2] = newCard()
	// above line give error because size is only 2

	//card_array = append(card_array, newCard())
	// above line give error because this is not slice
	fmt.Println("array_after_add", card_array)

}

func newCard() string {
	return "mycard"
}
