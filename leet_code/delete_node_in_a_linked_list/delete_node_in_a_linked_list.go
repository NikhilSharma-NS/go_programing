package delete_node_in_a_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteNode(node *ListNode) {

	*node = *node.Next

}

func PrintData(node *ListNode) {
	temp := node
	for temp != nil {
		fmt.Printf("[%v]", temp.Val)
		temp = temp.Next
	}
}
