package rotate_list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func RotateRight(head *ListNode, k int) *ListNode {
	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}

	current := head
	counter := 0
	nxtCurrent := &ListNode{}

	if k == 0 || head == nil || count == 0 || count == 1 || count == k {
		return head
	} else if k > count {
		k = k % count
		if k == 0 {
			return head
		}
	}

	for current != nil {
		counter++
		nxtCurrent = current.Next
		if counter == count-k {
			current.Next = nil
			tempnxtCurrent := nxtCurrent
			for tempnxtCurrent.Next != nil {
				tempnxtCurrent = tempnxtCurrent.Next
			}
			tempnxtCurrent.Next = head
			break
		}
		current = current.Next
	}

	return nxtCurrent
}

func PrintData(node *ListNode) {
	temp := node
	for temp != nil {
		fmt.Printf("[%v]", temp.Val)
		temp = temp.Next
	}
}
