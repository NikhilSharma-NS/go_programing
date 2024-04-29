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
	queue.Enqueue(nil)

	for len(queue.List) > 1 {
		currentNode := queue.Dequeue()
		if currentNode == nil {
			queue.Enqueue(nil)
			continue
		}
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

}
