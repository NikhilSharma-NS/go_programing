package main

type Node[T any] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func main() {

	left := &Node[int]{data: 10}
	right := &Node[int]{data: 12}

	root := Node[int]{data: 5}
	root.left = left
	root.right = right

}
