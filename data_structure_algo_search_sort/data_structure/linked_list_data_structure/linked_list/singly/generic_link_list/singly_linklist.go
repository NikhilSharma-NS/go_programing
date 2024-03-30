package main

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type User struct {
	name string
}

type SinglyList[T any] struct {
	head   *Node[T]
	length int
}

func (s *SinglyList[T]) insertAtEnd(value T) {

	node := &Node[T]{data: value}

	if s.head == nil {
		s.head = node
		s.length++
	} else {
		pre := s.head
		for pre.next != nil {
			pre = pre.next
		}
		pre.next = node

	}

}

func (s *SinglyList[T]) printList() {
	pre := s.head
	for pre != nil {
		fmt.Println(pre.data)
		pre = pre.next
	}
}

func main() {

	//n := Node[int]{data: 1}

	s := SinglyList[int]{}
	s.insertAtEnd(1)
	s.insertAtEnd(4)
	s.printList()

	s1 := SinglyList[string]{}
	s1.insertAtEnd("one")
	s1.insertAtEnd("four")
	s1.printList()

	s2 := SinglyList[User]{}
	s2.insertAtEnd(User{name: "jay"})
	s2.insertAtEnd(User{name: "vijay"})
	s2.printList()

}
