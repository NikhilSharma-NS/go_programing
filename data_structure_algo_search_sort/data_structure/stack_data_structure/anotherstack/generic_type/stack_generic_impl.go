package main

import "fmt"

type GenericType interface{}

type StackGenric struct {
	Gtypes []GenericType
}

func (s *StackGenric) pushDataG(value GenericType) {

	s.Gtypes = append(s.Gtypes, value)
}

func (s *StackGenric) display() []GenericType {

	return s.Gtypes
}

func (s *StackGenric) popData() (g *GenericType, p bool) {
	len := len(s.Gtypes)
	if len == 0 {
		return nil, false
	}
	item := s.Gtypes[len-1]
	s.Gtypes = s.Gtypes[:len-1]
	return &item, true
}

func main() {

	stack := StackGenric{Gtypes: nil}

	stack.pushDataG("key")
	stack.pushDataG("key1")

	fmt.Println(stack.display())

	fmt.Println(stack.popData())
	fmt.Println(stack.popData())
	fmt.Println(stack.popData())
	fmt.Println(stack.display())
}
