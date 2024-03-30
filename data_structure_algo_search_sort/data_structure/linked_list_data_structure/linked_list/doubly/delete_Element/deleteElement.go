package main

import "fmt"

type T1 interface {
	comparable
}

type Node[T T1] struct {
	prev *Node[T]
	next *Node[T]
	data T
}

func (head *Node[T]) deleteHead() *Node[T] {

	if head == nil && head.next == nil {
		return nil
	}

	pre := head

	head = head.next

	head.prev = nil
	pre.next = nil

	return head

}

func DisplayDoblyLinkList[T T1](head *Node[T]) {

	for head != nil {

		fmt.Printf("[%v]", head.data)
		head = head.next
	}
}

func main() {

	n := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3}}}
	DisplayDoblyLinkList[int](n)
	fmt.Println()
	n = n.deleteHead()
	DisplayDoblyLinkList[int](n)

}
