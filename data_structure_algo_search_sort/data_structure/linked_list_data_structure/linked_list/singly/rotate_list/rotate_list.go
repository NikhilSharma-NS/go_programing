package main

import (
	"fmt"
)

type T1 interface {
	comparable
}

type Node[T T1] struct {
	Data T
	Next *Node[T]
}

func (head *Node[T]) Rotate_listByKth(k int) *Node[T] {

	temp := head
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}

	current := head
	counter := 0
	nxtCurrent := &Node[T]{}

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

func (n *Node[T]) PrintData() {
	temp := n
	for temp != nil {
		fmt.Printf("[%v]", temp.Data)
		temp = temp.Next
	}
}

func main() {
	fmt.Println("Before Rotaion of [1,2,3,4,5]: by [2]th")
	node3 := &Node[int]{Data: 1, Next: &Node[int]{Data: 2, Next: &Node[int]{Data: 3, Next: &Node[int]{Data: 4, Next: &Node[int]{Data: 5}}}}}

	node3 = node3.Rotate_listByKth(2)
	node3.PrintData()
}
