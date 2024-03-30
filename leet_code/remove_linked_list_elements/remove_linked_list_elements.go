package remove_linked_list_elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, current := dummy, head
	for current != nil {
		temp := current.Next
		if current.Val == val {
			pre.Next = temp
		} else {
			pre = current
		}
		current = current.Next

	}
	return dummy.Next

}
