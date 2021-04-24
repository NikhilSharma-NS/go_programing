package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int64
}

func main() {
	root := &TreeNode{value: 1}
	root.left = &TreeNode{value: 2}
	root.right = &TreeNode{value: 3}
	root.left.left = &TreeNode{value: 4}
	root.left.right = &TreeNode{value: 5}
	root.right.left = &TreeNode{value: 6}
	root.right.right = &TreeNode{value: 7}
	root.preOrderTraversal()
}

func (treeNode *TreeNode) preOrderTraversal() {
	if treeNode != nil {
		fmt.Println(treeNode.value)
		treeNode.left.preOrderTraversal()
		treeNode.right.preOrderTraversal()
	}

}
