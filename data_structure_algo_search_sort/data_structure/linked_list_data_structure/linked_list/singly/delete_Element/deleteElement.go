package main

import (
	"fmt"
)

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

func (head *Node[T]) deleteElementAtKth(k int) *Node[T] {
	temp := head
	var pre *Node[T]

	// to handle empty
	if temp == nil {
		return nil
	}
	// to handle head
	if k == 1 {
		head = head.next
		return head
	}
	// to handle kth
	counter := 0
	for temp != nil {
		counter++
		if counter == k {
			pre.next = pre.next.next
		}
		pre = temp
		temp = temp.next
	}
	return head

}

func (head *Node[T]) deleteElement(k T) *Node[T] {
	temp := head
	var pre *Node[T]

	// to handle empty
	if temp == nil {
		return nil
	}
	// to handle head
	if k == head.data {
		head = head.next
		return head
	}
	// to handle kth
	for temp != nil {
		if temp.data == k {
			pre.next = pre.next.next
		}
		pre = temp
		temp = temp.next
	}
	return head

}

func (s *SNode[T]) deleteElementAtKth(k int) {

	temp := s.head

	if temp == nil {
		return
	} else if k == 1 {
		s.head = s.head.next
		return
	}

	count := 0
	pre := &Node[T]{}
	for temp != nil {
		count++
		if count == k {
			pre.next = pre.next.next
		}
		pre = temp
		temp = temp.next
	}

}

func (s *SNode[T]) deleteElement(k T) {

	temp := s.head

	if temp == nil {
		return
	} else if k == temp.data {
		s.head = s.head.next
		return
	}

	pre := &Node[T]{}
	for temp != nil {
		if k == temp.data {
			pre.next = pre.next.next
		}
		pre = temp
		temp = temp.next
	}

}

func (head *Node[T]) deleteElementAtEnd() *Node[T] {

	temp := head

	if temp == nil || temp.next == nil {
		return nil
	}

	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	return head

}

func (s *SNode[T]) deleteElementAtEnd() {
	temp := s.head

	if temp == nil || temp.next == nil {
		return
	}

	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
}

func (h *Node[T]) deleteHead() *Node[T] {
	//temp := h
	h = h.next
	return h
}

func (s *SNode[T]) deleteHead() {

	s.head = s.head.next
}

func (head *Node[T]) deleteElementAllCase(k T) *Node[T] {

	dummy := &Node[T]{next: head}

	pre, current := dummy, head

	for current != nil {
		temp := current.next
		if current.data == k {
			pre.next = temp
		} else {
			pre = current
		}
		current = current.next
	}
	return dummy.next

}

func (head *Node[T]) displayNodeData() {
	temp := head
	for temp != nil {
		fmt.Printf("[%v]", temp.data)
		temp = temp.next
	}
	fmt.Println("")
}
func main() {

	n := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3, next: &Node[int]{data: 4, next: &Node[int]{data: 5, next: &Node[int]{data: 6, next: &Node[int]{data: 7}}}}}}}

	//n.deleteElement(2)
	n.displayNodeData()
	fmt.Println("Deleting At Head")
	n = n.deleteHead()
	n.displayNodeData()
	fmt.Println("----------------------------")
	n.displayNodeData()
	fmt.Println("Deleting At End")
	n = n.deleteElementAtEnd()
	n.displayNodeData()
	fmt.Println("----------------------------")
	n.displayNodeData()
	fmt.Println("Deleting At kth(1) End")
	n = n.deleteElementAtKth(1)
	n.displayNodeData()
	fmt.Println("----------------------------")
	n.displayNodeData()
	fmt.Println("Deleting At kth(2) End")
	n = n.deleteElementAtKth(2)
	n.displayNodeData()
	fmt.Println("----------------------------")
	n.displayNodeData()
	n = n.deleteElement(5)
	fmt.Println("Deleting value ")
	n.displayNodeData()

	fmt.Println("***************************")
	n1 := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 3, next: &Node[int]{data: 4, next: &Node[int]{data: 5, next: &Node[int]{data: 6, next: &Node[int]{data: 7}}}}}}}
	s := SNode[int]{head: n1}
	fmt.Println("----------------------------")
	s.head.displayNodeData()
	fmt.Println("Deleting At Head")
	s.deleteHead()
	s.head.displayNodeData()
	fmt.Println("----------------------------")
	s.head.displayNodeData()
	fmt.Println("Deleting At End")
	s.head.deleteElementAtEnd()
	s.head.displayNodeData()
	fmt.Println("----------------------------")
	s.head.displayNodeData()
	fmt.Println("Deleting At kth(1) end")
	s.deleteElementAtKth(1)
	s.head.displayNodeData()

	fmt.Println("----------------------------")
	s.head.displayNodeData()
	fmt.Println("Deleting At kth(2) end")
	s.deleteElementAtKth(2)
	s.head.displayNodeData()

	n5 := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 6, next: &Node[int]{data: 3, next: &Node[int]{data: 4, next: &Node[int]{data: 5, next: &Node[int]{data: 6}}}}}}}

	n5.displayNodeData()
	fmt.Println("Deleting 6 from list")
	n5 = n5.deleteElementAllCase(6)
	n5.displayNodeData()
	fmt.Println("----------------------------")
	n6 := &Node[int]{}
	n6.displayNodeData()
	fmt.Println("Deleting 7 from Empty List")
	n6 = n6.deleteElementAllCase(7)
	n6.displayNodeData()
	fmt.Println("----------------------------")
	n7 := &Node[int]{data: 7, next: &Node[int]{data: 7, next: &Node[int]{data: 7, next: &Node[int]{data: 7}}}}
	n7.displayNodeData()
	fmt.Println("Deleting 7 from list")
	n7 = n7.deleteElementAllCase(7)
	n7.displayNodeData()

	n8 := &Node[int]{data: 1, next: &Node[int]{data: 2, next: &Node[int]{data: 2, next: &Node[int]{data: 1}}}}
	n8.displayNodeData()
	fmt.Println("Deleting 2 from list")
	n8 = n8.deleteElementAllCase(2)
	n8.displayNodeData()
}
