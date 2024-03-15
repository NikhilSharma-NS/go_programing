package main

import "fmt"

type Queuee[T any] struct {
	Items []T
}

func (q *Queuee[T]) Enqueue(v T) {
	q.Items = append(q.Items, v)
}

func (q *Queuee[T]) Dequeue() {
	if len(q.Items) > 0 {
		q.Items = q.Items[1:len(q.Items)]
	}

}

func (q *Queuee[T]) printQueue() {
	fmt.Println(q.Items)
}
func main() {
	q := Queuee[int]{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.printQueue()
	q.Dequeue()
	q.printQueue()

}
