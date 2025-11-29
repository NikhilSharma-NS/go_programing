package data_type

import (
	"fmt"
)

func StringDataType() {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World India"
	fmt.Println()
	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}
