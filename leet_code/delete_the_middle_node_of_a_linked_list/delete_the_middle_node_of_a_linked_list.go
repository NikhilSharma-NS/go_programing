package delete_the_middle_node_of_a_linked_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteMiddle(head *ListNode) *ListNode {
	count := 0
	mid := 0
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}

	mid = (count / 2)

	dummy := &ListNode{Next: head}
	pre, curr := dummy, head
	counter := 0
	for curr != nil {
		next := curr.Next
		if counter == mid {
			pre.Next = next
		} else {
			pre = curr
		}
		curr = curr.Next
		counter++
	}

	return dummy.Next
}

func PrintData(node *ListNode) {
	temp := node
	for temp != nil {
		fmt.Printf("[%v]", temp.Val)
		temp = temp.Next
	}
}
