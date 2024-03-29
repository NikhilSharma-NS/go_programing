package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

func (head *Node[T]) traversal() {
	temp := head
	for temp != nil {
		fmt.Printf("[%v]", temp.data)
		temp = temp.next
	}
}

func main() {
	n := Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3, next: &Node[int]{data: 4}}}}
	n.traversal()
}
