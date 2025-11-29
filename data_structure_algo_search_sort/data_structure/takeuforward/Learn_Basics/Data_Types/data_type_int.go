package data_type

import (
	"fmt"
)

func IntDataType() {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)
}
