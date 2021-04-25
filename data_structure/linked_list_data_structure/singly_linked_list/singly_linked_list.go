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

func (singlyLinkedList *SinglyLinkedList) insert(val int) {
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

func (singlyLinkedList *SinglyLinkedList) Traverse() {
	if singlyLinkedList.head == nil {
		fmt.Println("Empty List")
		return
	}
	current := singlyLinkedList.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}

}

func main() {

	singlyLinkedLis := &SinglyLinkedList{}
	singlyLinkedLis.insert(2)
	singlyLinkedLis.insert(3)
	singlyLinkedLis.insert(4)
	singlyLinkedLis.Traverse()

}
