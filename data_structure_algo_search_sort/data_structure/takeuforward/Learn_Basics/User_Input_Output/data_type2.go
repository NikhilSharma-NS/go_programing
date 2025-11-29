package input_output

import (
	"fmt"
)

func TakeInput2() {
	var i, j int

	fmt.Print("Type two numbers: ")
	fmt.Scanf("%v %v", &i, &j)
	fmt.Println("Your numbers are:", i, "and", j)
}
