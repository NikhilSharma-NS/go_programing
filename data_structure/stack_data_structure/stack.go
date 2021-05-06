package main

import "fmt"

type Stack []string

func (stack *Stack) size() int {
	return len(*stack)
}

func (stack *Stack) isEmpty() bool {
	return len(*stack) == 0
}

func (stack *Stack) Push(str string) {
	*stack = append(*stack, str)
}

// remove and return the top element of stack , return false if stack is empty
func (stack *Stack) Pop() (string, bool) {

	if stack.isEmpty() {
		return "", false
	} else {
		index := stack.size() - 1  // Get the index of the top most element.
		element := (*stack)[index] // Index into the slice and obtain the element.
		*stack = (*stack)[:index]  // Remove it from the stack by slicing it off.
		return element, true
	}
}

func main() {
	stack := Stack{}
	fmt.Println("size---", stack.size())
	stack.Push("a")
	stack.Push("b")
	fmt.Println("size---", stack.size())
	for len(stack) != 0 {
		value, isdeleted := stack.Pop()
		if isdeleted == true {
			fmt.Println(value)
		}
	}
	fmt.Println("size---", stack.size())

}
