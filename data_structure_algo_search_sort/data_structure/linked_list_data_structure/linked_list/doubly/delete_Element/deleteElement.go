package main

import (
	"fmt"
)

type T1 interface {
	comparable
}

type Node[T T1] struct {
	prev *Node[T]
	next *Node[T]
	data T
}

func (head *Node[T]) deleteHead() *Node[T] {
	if head == nil && head.next == nil {
		return nil
	}
	pre := head
	head = head.next
	head.prev = nil
	pre.next = nil
	return head
}

func (head *Node[T]) deleteTail() *Node[T] {
	if head == nil && head.next == nil {
		return nil
	}
	tail := head
	for tail.next != nil {
		tail = tail.next
	}
	n := tail.prev
	n.next = nil
	tail.prev = nil

	return head
}

func (head *Node[T]) deletekthNode(k int) *Node[T] {

	if head == nil {
		return nil
	}
	counter := 0
	Knode := head

	for Knode != nil {
		counter++
		if counter == k {
			break
		}
		Knode = Knode.next
	}

	pre := Knode.prev
	next := Knode.next

	if pre == nil && next == nil {
		return nil
	} else if pre == nil {
		return head.deleteHead()
	} else if next == nil {
		return head.deleteTail()
	} else {
		pre.next = next
		next.prev = pre
	}

	Knode.next = nil
	Knode.prev = nil
	return head
}

func (head *Node[T]) deleteValueNode(n *Node[T]) {

	if head == nil {
		return
	}

	pre := n.prev
	next := n.next

	if next == nil {
		pre.next = nil
		n.prev = nil
		return
	}

	pre.next = next
	next.prev = pre

	n.next = nil
	n.prev = nil

}

func DisplayDoblyLinkList[T T1](head *Node[T]) {
	for head != nil {
		fmt.Printf("[%v]", head.data)
		head = head.next
	}
}

func CovertArrayToDoblyLinkList[T T1](arr []T) (head *Node[T]) {

	h := &Node[T]{data: arr[0]}
	pre := h

	for i := 1; i < len(arr); i++ {
		temp := &Node[T]{data: arr[i], prev: pre, next: nil}
		pre.next = temp
		pre = pre.next
		//pre = temp

	}
	return h

}
func main() {

	n := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3}}}
	DisplayDoblyLinkList[int](n)
	fmt.Println()
	n = n.deleteHead()
	DisplayDoblyLinkList[int](n)
	fmt.Println()
	arr := []int{1, 2, 3, 4, 5}
	n1 := &Node[int]{}
	n1 = CovertArrayToDoblyLinkList(arr)
	DisplayDoblyLinkList[int](n1)
	fmt.Println()
	n1 = n1.deletekthNode(1)
	DisplayDoblyLinkList[int](n1)
	fmt.Println("By value")
	n1.deleteValueNode(n1.next.next)
	DisplayDoblyLinkList[int](n1)
}
