package tasks

import (
	"fmt"
)

var stack []string

func StackprintLIFO() {

	stack = append(stack, "first")

	stack = append(stack, "second")

	for len(stack) > 0 {
		length := len(stack) - 1
		fmt.Println(stack[length])
		stack = stack[:length]
	}

}
