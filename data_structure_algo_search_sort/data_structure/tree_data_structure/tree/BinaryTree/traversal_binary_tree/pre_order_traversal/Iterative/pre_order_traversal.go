package main

import "fmt"

type TreeNode[T any] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

// type T2 interface {
// 	comparable
// }

type Stack[T1 any] struct {
	Items []T1
}

func (s *Stack[T1]) PushItem(v T1) {
	s.Items = append(s.Items, v)
}

func (s *Stack[T1]) PopItem() {
	length := len(s.Items)
	if len(s.Items) > 0 {
		s.Items = s.Items[:length-1]
	}
}

func (s *Stack[T1]) PeekItem() (v T1) {
	if len(s.Items) == 0 {
		return v
	}

	return s.Items[len(s.Items)-1]
}

func (s Stack[T1]) printItems() {
	fmt.Println(s.Items)
}

func (root *TreeNode[T]) PreOrderTraversal() {

	pre_order_iterative := []T{}
	stack := &Stack[TreeNode[T]]{}

	stack.PushItem(*root)

	for len(stack.Items) != 0 {
		top := stack.PeekItem()
		stack.PopItem()
		pre_order_iterative = append(pre_order_iterative, top.data)

		if top.right != nil {
			stack.PushItem(*top.right)
		}
		if top.left != nil {
			stack.PushItem(*top.left)
		}

	}
	fmt.Println(pre_order_iterative)

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
