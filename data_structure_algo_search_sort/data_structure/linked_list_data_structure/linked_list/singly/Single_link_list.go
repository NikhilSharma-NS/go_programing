package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type SList struct {
	head   *Node
	length int
}

func (s *SList) insertAtEnd(value int) {
	head := s.head
	n := &Node{val: value}
	if head == nil {
		s.head = n
		return
	}
	for head.next != nil {
		head = head.next
	}
	head.next = n
	s.length = s.length + 1

}

func (s SList) displayAllData() {
	fmt.Println("Displaying Records")
	head := s.head

	for head != nil {
		fmt.Printf("[%v]", head.val)
		fmt.Print("<-->")
		head = head.next
	}
	fmt.Println()
}

func main() {
	s := &SList{head: &Node{val: 1, next: &Node{val: 2}}, length: 2}
	s.insertAtEnd(3)
	s.insertAtEnd(4)
	s.displayAllData()

	s1 := &SList{}
	s1.insertAtEnd(3)
	s1.insertAtEnd(4)
	s1.displayAllData()

}
