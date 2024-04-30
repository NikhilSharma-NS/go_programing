package main

import "fmt"

type TreeNode[T any] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

type Queue[T any] struct {
	List [](*TreeNode[T])
}

// LevelOrderTraversal(BFS)
func (node *TreeNode[T]) LevelOrderTraversal() {
	levelorder := []T{}
	queue := Queue[T]{}

	if node == nil {
		return
	}

	queue.Enqueue(node)

	for len(queue.List) > 0 {
		currentNode := queue.Dequeue()
		levelorder = append(levelorder, currentNode.data)
		if currentNode.left != nil {
			queue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			queue.Enqueue(currentNode.right)

		}

	}

	fmt.Println(levelorder)

}

func (q *Queue[T]) Enqueue(element *TreeNode[T]) {

	q.List = append(q.List, element)
}

func (q *Queue[T]) Dequeue() *TreeNode[T] {

	if len(q.List) == 0 {
		return nil
	}
	element := q.List[0]
	q.List = q.List[1:]
	return element

}

func main() {
	root := TreeNode[int]{data: 1}

	left := &TreeNode[int]{data: 2}
	left_left := &TreeNode[int]{data: 4}
	left_right := &TreeNode[int]{data: 5}
	left.left = left_left
	left.right = left_right

	right := &TreeNode[int]{data: 3}
	right_left := &TreeNode[int]{data: 6}
	right_right := &TreeNode[int]{data: 7}
	right.left = right_left
	right.right = right_right

	root.left = left
	root.right = right

	root.LevelOrderTraversal()

	roots := TreeNode[string]{data: "data"}

	lefts := &TreeNode[string]{data: "papa"}
	left_lefts := &TreeNode[string]{data: "papa_son1"}
	left_rights := &TreeNode[string]{data: "papa_son2"}
	lefts.left = left_lefts
	lefts.right = left_rights

	rights := &TreeNode[string]{data: "uncle"}
	right_lefts := &TreeNode[string]{data: "uncle_son1"}
	right_rights := &TreeNode[string]{data: "uncle_son2"}
	rights.left = right_lefts
	rights.right = right_rights

	roots.left = lefts
	roots.right = rights

	roots.LevelOrderTraversal()

}
