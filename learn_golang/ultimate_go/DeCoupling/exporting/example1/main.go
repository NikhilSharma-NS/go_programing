package main

import (
	"fmt"

	"example2"
)

type AlterCounter int

type alterCounter int

func New(val int) alterCounter {
	return alterCounter(val)
}

func main() {

	c := New(10)

	fmt.Println(c)

	user := example2.User{}

	fmt.Println(user)

}

// Exporting provider ability to declare the indentifier is accessible to code aouside of the package it's declared In.
