package main

import (
	"fmt"
)

type T1 interface {
	comparable
}

type Node[T T1] struct {
	data T
	next *Node[T]
}

type Snode[T T1] struct {
	head *Node[T]
}

func (s *Snode[T]) serchElement(value T) bool {

	temp := s.head

	for temp != nil {
		if value == temp.data {
			return true
		}
		temp = temp.next
	}
	return false

}

func (head *Node[T]) serchElement(value T) bool {

	temp := head

	for temp != nil {
		if value == temp.data {
			return true
		}
		temp = temp.next
	}
	return false
}

func main() {

	n := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3, next: &Node[int]{data: 4}}}}

	s := &Snode[int]{head: n}

	found_or_not := n.serchElement(1)

	fmt.Println(found_or_not)

	found_or_not_s := s.serchElement(6)

	fmt.Println(found_or_not_s)

}
