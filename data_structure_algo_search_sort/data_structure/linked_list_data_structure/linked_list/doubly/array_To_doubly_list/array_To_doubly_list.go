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

func covertArrayToDoblyLinkList[T T1](arr []T) (head *Node[T]) {

	h := &Node[T]{data: arr[0]}
	pre := h

	for i := 1; i < len(arr); i++ {
		temp := &Node[T]{data: arr[i], prev: pre, next: nil}
		pre.next = temp
		pre = pre.next
		//pre = temp

	}
	return h

}

func DisplayDoblyLinkList[T T1](head *Node[T]) {

	for head != nil {

		fmt.Printf("[%v]", head.data)
		head = head.next
	}
}

func main() {

	arr := []int{1, 2, 3, 4, 5}

	n := &Node[int]{}
	n = covertArrayToDoblyLinkList(arr)
	DisplayDoblyLinkList[int](n)

}
