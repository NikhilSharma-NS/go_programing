package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int64
}

func main() {
        results:= []int64{}
	root := &TreeNode{value: 1}
	root.left = &TreeNode{value: 2}
	root.right = &TreeNode{value: 3}
	root.left.left = &TreeNode{value: 4}
	root.left.right = &TreeNode{value: 5}
	root.right.left = &TreeNode{value: 6}
	root.right.right = &TreeNode{value: 7}
	root.preOrderTraversal(&results)
	fmt.Println(results)
	
}

func (treeNode *TreeNode) preOrderTraversal(result *[]int64) {
	if treeNode != nil {
		*result =append(*result,treeNode.value)
		treeNode.left.preOrderTraversal(result)
		treeNode.right.preOrderTraversal(result)
	}

}
