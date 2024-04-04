package add_two_numbers

import "fmt"

/**
 * Definition for singly-linked list.

 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l2 == nil && l1 != nil {
		return l1
	}
	sum := l1.Val + l2.Val
	next := AddTwoNumbers(l1.Next, l2.Next)

	if sum >= 10 {
		sum %= 10
		next = AddTwoNumbers(next, &ListNode{
			Val: 1,
		})
	}
	return &ListNode{
		Val:  sum,
		Next: next,
	}

}

func (l *ListNode) PrintData() {
	for l != nil {
		fmt.Printf("[%v]", l.Val)
		l = l.Next
	}

}
