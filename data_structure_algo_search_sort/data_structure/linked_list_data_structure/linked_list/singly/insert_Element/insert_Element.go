package main

import "fmt"

type T1 interface {
	comparable
}

type Node[T T1] struct {
	data T
	next *Node[T]
}

type SNode[T T1] struct {
	head *Node[T]
}

func (head *Node[T]) insertAtHead(value T) *Node[T] {

	// head = &Node[T]{data: value, next: head}

	// return head
	return &Node[T]{data: value, next: head}

}

func (head *Node[T]) insertAtLast(value T) *Node[T] {

	temp := head
	n := &Node[T]{data: value}

	if temp == nil {
		head = n
	}

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = n

	return head

}

func (head *Node[T]) insertAtkthPosition(value T, index int) *Node[T] {

	newNode := &Node[T]{data: value}

	if head == nil {
		if index == 1 {
			return newNode
		} else {
			return head
		}
	}

	if index == 1 {
		newNode.next = head
		return newNode
	}

	counter := 0
	current := head

	for current != nil {
		counter++
		//	nxtCurrent := current.next
		if counter == index-1 {
			newNode.next = current.next
			current.next = newNode
			break
		}
		current = current.next
	}
	return head
}

func (head *Node[T]) insertBeforeValue(value_node T, val T) *Node[T] {

	newNode := &Node[T]{data: value_node}

	if head == nil {
		return nil
	}

	if head.data == val {
		newNode.next = head
		return newNode
	}
	current := head

	for current.next != nil {
		//	nxtCurrent := current.next
		if val == current.next.data {
			newNode.next = current.next
			current.next = newNode
			break
		}
		current = current.next
	}
	return head
}

func (sNode *SNode[T]) insertAtHead(value T) {

	sNode.head = &Node[T]{data: value, next: sNode.head}

}

func (sNode *SNode[T]) insertAtLast(value T) {

	temp := sNode.head
	n := &Node[T]{data: value}

	if temp == nil {
		sNode.head = n
	}

	for temp.next != nil {
		temp = temp.next
	}
	temp.next = n

}
func (head *Node[T]) printData() {
	temp := head

	for temp != nil {
		fmt.Printf("[%v]", temp.data)
		temp = temp.next
	}

}

func (sNode *SNode[T]) insertAtkthPosition(value T, index int) {
	newNode := &Node[T]{data: value}
	if sNode.head == nil {
		if index == 1 {
			sNode.head = newNode
		}
	}

	if index == 1 {
		newNode.next = sNode.head
		sNode.head = newNode
	}

	current := sNode.head
	counter := 0
	for current != nil {
		counter++
		if counter == index-1 {
			newNode.next = current.next
			current.next = newNode
			break
		}
		current = current.next
	}

}
func (sNode *SNode[T]) insertBeforeValue(value_node T, val T) {

	newNode := &Node[T]{data: value_node}

	if sNode.head == nil {
		return
	}

	if sNode.head.data == val {
		newNode.next = sNode.head
		return
	}
	current := sNode.head

	for current.next != nil {
		//	nxtCurrent := current.next
		if val == current.next.data {
			newNode.next = current.next
			current.next = newNode
			break
		}
		current = current.next
	}
}

func (head *SNode[T]) printData() {
	temp := head.head

	for temp != nil {
		fmt.Printf("[%v]", temp.data)
		temp = temp.next
	}

}

func main() {

	n := &Node[int]{data: 5}
	n = n.insertAtHead(2)
	n = n.insertAtHead(3)
	n = n.insertAtLast(6)
	n = n.insertAtkthPosition(10, 2)
	n = n.insertAtkthPosition(11, 1)
	n = n.insertAtkthPosition(12, 4)
	n = n.insertBeforeValue(19, 3)
	n.printData()

	fmt.Println()
	nx1 := &Node[int]{data: 5}
	sNode := &SNode[int]{head: nx1}
	sNode.insertAtHead(2)
	sNode.insertAtHead(3)
	sNode.insertAtLast(6)
	sNode.insertAtkthPosition(10, 2)
	sNode.insertAtkthPosition(11, 1)
	sNode.insertAtkthPosition(12, 4)
	sNode.insertBeforeValue(19, 3)
	sNode.printData()
}
