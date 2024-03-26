package main

import "fmt"

const SIZE = 8

type Node struct {
	Left  *Node
	Right *Node
	Val   string
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type HasMap map[string]*Node

type Cache struct {
	Queue  Queue
	HasMap HasMap
}

func (c *Cache) Check(s string) {

	node := &Node{}

	if existing_node, isExists := c.HasMap[s]; isExists {
		node = c.Remove(existing_node)
	} else {
		node = &Node{Val: s}
	}

	c.Add(node)
	c.HasMap[s] = node

}

func (c *Cache) Remove(node *Node) *Node {

	right := node.Right
	left := node.Left

	right.Left = left
	left.Right = right

	c.Queue.Length -= 1

	delete(c.HasMap, node.Val)

	return node

}

func (c *Cache) Add(node *Node) {

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = node
	node.Left = c.Queue.Head
	node.Right = tmp
	tmp.Left = node

	c.Queue.Length += 1

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Display() {

	c.Queue.Display()
}

func (q Queue) Display() {

	node := q.Head.Right
	fmt.Printf("Len[%d :", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("[%s]", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func NewCache() Cache {

	return Cache{Queue: NewQueue(), HasMap: HasMap{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	head.Left = head
	return Queue{Head: head, Tail: tail}
}

func main() {

	cache := NewCache()
	for _, val := range []string{"A", "B", "C", "D", "E", "F"} {
		cache.Check(val)
		cache.Display()
	}

	cache.Check("G")
	cache.Display()

}
