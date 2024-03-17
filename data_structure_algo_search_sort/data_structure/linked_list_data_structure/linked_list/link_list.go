package main

import "fmt"

type nodeK struct {
	data int
	next *nodeK
}

type linklist struct {
	head   *nodeK
	length int
}

func (l *linklist) insertAtFront(n *nodeK) {
	head := l.head
	l.head = n
	l.head.next = head
	l.length = l.length + 1

}



func (l *linklist) printData() {
	h := l.head
	for l.length != 0 {
		fmt.Printf("List : %d \n", h.data)
		h = h.next
		l.length--
	}
	fmt.Println()
}

func main() {

	mylist := linklist{}
	mylist.insertAtFront(&nodeK{data: 1})
	mylist.insertAtFront(&nodeK{data: 2})
	mylist.insertAtFront(&nodeK{data: 3})
	mylist.insertAtFront(&nodeK{data: 4})
	mylist.printData()
}
