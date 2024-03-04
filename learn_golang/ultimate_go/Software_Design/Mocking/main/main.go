package main

import (
	pub "Gorepo/go_programing/learn_golang/ultimate_go/Software_Design/Mocking"
	"fmt"
)

type Publisher interface {
	Publish()
	Subscrib()
}

type mock struct{}

func (m *mock) Publish() {
	fmt.Println("testing Publish")
}

func (m *mock) Subscrib() {
	fmt.Println("testing Subscrib")
}

func main() {
	pub := []Publisher{
		pub.New("lo"), &mock{},
	}

	for _, p := range pub {
		p.Publish()
		p.Subscrib()
	}
}
