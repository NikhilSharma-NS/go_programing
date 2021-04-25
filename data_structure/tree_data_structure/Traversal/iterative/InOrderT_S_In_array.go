package main

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
}

func main() {
	root := &TreeNode{value: 1}
	root.left = &TreeNode{value: 2}
	root.right = &TreeNode{value: 3}
	root.left.left = &TreeNode{value: 4}
	root.left.right = &TreeNode{value: 5}
	root.right.left = &TreeNode{value: 6}
	root.right.right = &TreeNode{value: 7}
	results := root.inOrderTraversal()
	fmt.Println(results)

}

func (root *TreeNode) inOrderTraversal() []int {
	result := []int{}
	if root == nil {
		return result
	}
	current := root
	stack := arraystack.New()
	for stack.Size() > 0 || current != nil {
		if current != nil {
			stack.Push(current)
			current = current.left
		} else {
			inferfaceObject, _ := stack.Pop()
			current = inferfaceObject.(*TreeNode)
			result = append(result, current.value)
			current = current.right
		}
	}
	return result

}
