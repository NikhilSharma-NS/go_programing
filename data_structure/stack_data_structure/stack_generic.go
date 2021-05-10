package main

import "fmt"

type ItemType interface{}

type Stack_Gen struct {
	items []ItemType
}

func (stack_gen *Stack_Gen) PushItem(itemtype ItemType) {

	if itemtype == nil {
		stack_gen.items = []ItemType{}
	}
	stack_gen.items = append(stack_gen.items, itemtype)

}

func (stack_gen *Stack_Gen) Pop() *ItemType {
	if len(stack_gen.items) == 0 {
		return nil
	}
	// Popping item from items slice.
	item := stack_gen.items[len(stack_gen.items)-1]
	//Adjusting the item's length accordingly
	stack_gen.items = stack_gen.items[:len(stack_gen.items)-1]

	return &item
}

func (stack_gen *Stack_Gen) GetAll() []ItemType {

	return stack_gen.items
}

func main() {
	stackgen := Stack_Gen{items: nil}
	stackgen.PushItem("jay")
	stackgen.PushItem("10")
	stackgen.Pop()
	stackgen.PushItem("20")
	stackgen.PushItem("30")
	stackgen.PushItem("40")
	fmt.Println(stackgen.GetAll())
	stackgen.Pop()
	fmt.Println(stackgen.GetAll())
}
