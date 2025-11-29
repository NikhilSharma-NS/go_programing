package input_output

import (
	"fmt"
)

func TakeInput3() {
	var i, j int

	fmt.Print("Type two numbers: ")
	fmt.Scanln(&i, &j)
	fmt.Println("Your numbers are:", i, "and", j)
}
