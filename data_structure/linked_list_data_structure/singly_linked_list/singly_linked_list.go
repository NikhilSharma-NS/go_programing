package main

import (
	"fmt"
)

type Node struct {
	next  *Node
	value int
}

type SinglyLinkedList struct {
	head   *Node
	length int
}

func (singlyLinkedList *SinglyLinkedList) insertAtBack(val int) {
	node := &Node{value: val}
	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		current := singlyLinkedList.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	singlyLinkedList.length++
	return
}

func (singlyLinkedList *SinglyLinkedList) insertAtFront(val int) {

	node := &Node{value: val}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		node.next = singlyLinkedList.head
		singlyLinkedList.head = node
	}
	singlyLinkedList.length++

}

func (singlyLinkedList *SinglyLinkedList) deleteAtFront() error {

	if singlyLinkedList.head == nil {
		return fmt.Errorf("List is empty")
	}
	singlyLinkedList.head = singlyLinkedList.head.next
	singlyLinkedList.length--
	return nil
}

func (singlyLinkedList *SinglyLinkedList) deleteAtBack() error {
	if singlyLinkedList.head == nil {
		return fmt.Errorf("List is Empty")
	}
	current := singlyLinkedList.head
	var prev *Node
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
	} else {
		singlyLinkedList.head = nil
	}
	singlyLinkedList.length--
	return nil
}

func (singlyLinkedList *SinglyLinkedList)size(){
	return singlyLinkedList.length
}

func (singlyLinkedList *SinglyLinkedList) traverse() error {
	if singlyLinkedList.head == nil {
		return fmt.Errorf("List is Empty")
	}
	current := singlyLinkedList.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	return nil
}

func main() {

	singlyLinkedLis := &SinglyLinkedList{}
	singlyLinkedLis.insertAtBack(2)
	singlyLinkedLis.insertAtBack(3)
	singlyLinkedLis.insertAtBack(4)
	singlyLinkedLis.traverse()

}
