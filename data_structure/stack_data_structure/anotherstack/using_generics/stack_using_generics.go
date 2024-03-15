package main

import "fmt"

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) PushItem(v T) {
	s.Items = append(s.Items, v)
}

func (s *Stack[T]) PopItem() {
	length := len(s.Items)
	if len(s.Items) > 0 {
		s.Items = s.Items[:length-1]
	}
}

func (s Stack[T]) printItems() {
	fmt.Println(s.Items)
}

func main() {
	s := Stack[int]{}
	s.PushItem(1)
	s.PushItem(2)
	s.printItems()
	s.PopItem()
	s.printItems()
}
