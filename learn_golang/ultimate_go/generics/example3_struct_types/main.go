package main

import "fmt"

type node[T any] struct {
	Data T
	next *node[T]
	prev *node[T]
}

type list[T any] struct {
	first *node[T]
	last  *node[T]
}

func (l *list[T]) add(data T) *node[T] {
	n := node[T]{Data: data, prev: l.last}

	if l.first == nil {
		l.first = &n
		l.last = &n
		return &n
	}
	l.last.next = &n
	l.last = &n
	return &n

}

type user struct {
	name string
}

func main() {
	var lv list[user]

	n1 := lv.add(user{"nik"})
	n2 := lv.add(user{"jay"})

	fmt.Println(n1.Data, n2.Data)

	var lv1 list[*user]

	n3 := lv1.add(&user{"nik"})
	n4 := lv1.add(&user{"jay"})

	fmt.Println(n3.Data, n4.Data)

}
