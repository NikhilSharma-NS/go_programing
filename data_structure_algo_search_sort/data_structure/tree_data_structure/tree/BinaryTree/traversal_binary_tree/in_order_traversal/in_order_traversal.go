package main

import "fmt"

type Node[T any] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func (root *Node[T]) InOrderTraversal() {

	if root == nil {
		return
	}
	root.left.InOrderTraversal()
	fmt.Printf("[%v]", root.data)
	root.right.InOrderTraversal()

}

func main() {
	root := Node[int]{data: 1}

	left := &Node[int]{data: 2}
	left_left := &Node[int]{data: 4}
	left_right := &Node[int]{data: 5}
	left.left = left_left
	left.right = left_right

	right := &Node[int]{data: 3}
	right_left := &Node[int]{data: 6}
	right_right := &Node[int]{data: 7}
	right.left = right_left
	right.right = right_right

	root.left = left
	root.right = right

	root.InOrderTraversal()
}
