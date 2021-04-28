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

func (singlyLinkedList *SinglyLinkedList) insertAtIndex(index int, value int) {
	node := &Node{value: value}
	var prev *Node
	if index < 0 {
		return
	} else if index == 0 {
		singlyLinkedList.insertAtFront(value)
		return
	}
	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		n := singlyLinkedList.getAt(index)
		node.next = n
		prev = singlyLinkedList.getAt(index - 1)
		prev.next = node
	}
	singlyLinkedList.length++
}

func (singlyLinkedList *SinglyLinkedList) getAt(position int) *Node {
	current := singlyLinkedList.head

	if position < 0 {
		return current
	} else if position > singlyLinkedList.size()-1 {
		return nil
	} else {
		for counter := 0; counter < position; counter++ {
			current = current.next
		}
	}
	return current
}

func (singlyLinkedList *SinglyLinkedList) deleteAtFront() error {

	if singlyLinkedList.head == nil {
		return fmt.Errorf("List is empty")
	}
	singlyLinkedList.head = singlyLinkedList.head.next
	singlyLinkedList.length--
	return nil
}

func (singlyLinkedList *SinglyLinkedList) deleteNode(input int) {
	current := singlyLinkedList.head
	var pre *Node
	if current != nil && current.value == input {
		singlyLinkedList.head = current.next
		singlyLinkedList.length--
		return
	} else {
		for current != nil && current.value != input {
			pre = current
			current = current.next
		}
	}
	if current == nil {
		return
	}
	pre.next = current.next
	singlyLinkedList.length--
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

func (singlyLinkedList *SinglyLinkedList) size() int {
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
	singlyLinkedLis.insertAtBack(6)

	singlyLinkedLis.insertAtIndex(0, 5)
	singlyLinkedLis.deleteNode(4)
	singlyLinkedLis.traverse()

}
