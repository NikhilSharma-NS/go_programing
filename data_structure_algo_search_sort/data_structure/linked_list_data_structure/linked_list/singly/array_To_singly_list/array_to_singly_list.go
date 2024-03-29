package main

import (
	"fmt"
)

type Node[T any] struct {
	data T
	next *Node[T]
}

type SNode[T any] struct {
	head *Node[T]
	len  int
}

func covertArrayToLinkList[T any](arr []T) *SNode[T] {
	out := &SNode[T]{}
	for i := 0; i < len(arr); i++ {
		out.insertAtEnd(arr[i])
	}
	return out
}

func covertArrayToLinkListway2[T any](arr []T) *SNode[T] {
	out := &SNode[T]{}
	temp := &Node[T]{data: arr[0]}
	out.head = temp
	mover := out.head
	for i := 1; i < len(arr); i++ {
		t := &Node[T]{data: arr[i]}
		mover.next = t
		mover = t
	}
	return out
}

func covertArrayToLinkListway3[T any](arr []T) *Node[T] {
	head := &Node[T]{data: arr[0]}
	mover := head
	for i := 1; i < len(arr); i++ {
		temp := &Node[T]{data: arr[i]}
		mover.next = temp
		mover = temp
	}
	return head
}

func (s *SNode[T]) insertAtEnd(value T) {
	head := s.head
	n := &Node[T]{data: value}
	if head == nil {
		s.head = n
		return
	}
	for head.next != nil {
		head = head.next
	}
	head.next = n
	s.len = s.len + 1

}

func (s SNode[T]) displayAllData() {
	fmt.Println("Displaying Records")
	head := s.head

	for head != nil {
		fmt.Printf("[%v]", head.data)
		fmt.Print("<-->")
		head = head.next
	}
	fmt.Println()
}

func main() {

	s := covertArrayToLinkList([]int{1, 2, 3, 4})

	s.displayAllData()

	s1 := covertArrayToLinkListway2([]int{1, 2, 3, 4})

	s1.displayAllData()

	s2 := covertArrayToLinkListway2([]string{"1a", "2a", "3a", "4a"})

	s2.displayAllData()

	s2Node := covertArrayToLinkListway3([]string{"1ab", "2ab", "3ab", "4ab"})
	s3 := &SNode[string]{head: s2Node}
	s3.displayAllData()

}
