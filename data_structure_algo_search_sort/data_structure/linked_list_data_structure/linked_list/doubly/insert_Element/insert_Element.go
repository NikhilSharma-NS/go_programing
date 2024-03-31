package main

import "fmt"

type T1 interface {
	comparable
}

type Node[T T1] struct {
	prev *Node[T]
	next *Node[T]
	data T
}

func (head *Node[T]) insertBeforeHead(value T) *Node[T] {

	newHead := &Node[T]{data: value, next: head}

	head.prev = newHead

	return newHead

}

func (head *Node[T]) insertAtTail(value T) *Node[T] {
	neNode := &Node[T]{data: value}

	tail := head

	for tail.next != nil {
		tail = tail.next
	}

	tail.next = neNode
	neNode.prev = tail

	return head

}

func (head *Node[T]) insertBeforeTail(value T) *Node[T] {

	tail := head

	for tail.next != nil {
		tail = tail.next
	}
	pre := tail.prev
	neNode := &Node[T]{prev: pre, data: value, next: tail}
	pre.next = neNode
	tail.prev = neNode

	return head

}

func (head *Node[T]) inserteBeforekth(k int, value T) *Node[T] {

	if k == 1 {
		return head.insertBeforeHead(value)
	}
	temp := head
	counter := 0

	for temp != nil {
		counter++
		if k == counter {
			break
		}
		temp = temp.next
	}

	pre := temp.prev
	newNode := &Node[T]{prev: pre, data: value, next: temp}
	pre.next = newNode
	temp.prev = newNode
	return head

}

func (head *Node[T]) insertBeforeNode(n *Node[T], value T) {

	pre := n.prev
	newNode := &Node[T]{prev: pre, data: value, next: n}

	pre.next = newNode

	n.prev = newNode

}

func DisplayDoblyLinkList[T T1](head *Node[T]) {
	for head != nil {
		fmt.Printf("[%v]", head.data)
		head = head.next
	}
}

func main() {

	h := &Node[int]{data: 20}

	h = h.insertBeforeHead(10)

	DisplayDoblyLinkList[int](h)
	fmt.Println()
	h = h.insertAtTail(30)

	DisplayDoblyLinkList[int](h)

	h = h.insertBeforeTail(40)
	fmt.Println()
	DisplayDoblyLinkList[int](h)

	fmt.Println()

	h = h.inserteBeforekth(2, 50)

	DisplayDoblyLinkList[int](h)

	fmt.Println()

	h.insertBeforeNode(h.next.next, 70)
	DisplayDoblyLinkList[int](h)

}
