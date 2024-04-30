package main

import "fmt"

type TreeNode[T any] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func (root *TreeNode[T]) PreOrderTraversal() {

	if root == nil {
		return
	}

	fmt.Printf("[%v]", root.data)
	root.left.PreOrderTraversal()
	root.right.PreOrderTraversal()

}

func main() {
	root := TreeNode[int]{data: 1}

	left := &TreeNode[int]{data: 2}
	left_left := &TreeNode[int]{data: 4}
	left_right := &TreeNode[int]{data: 5}
	left.left = left_left
	left.right = left_right

	right := &TreeNode[int]{data: 3}
	right_left := &TreeNode[int]{data: 6}
	right_right := &TreeNode[int]{data: 7}
	right.left = right_left
	right.right = right_right

	root.left = left
	root.right = right

	root.PreOrderTraversal()
}
