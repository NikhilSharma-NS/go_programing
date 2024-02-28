package main

import "fmt"

func main() {
	fmt.Println("data")

	const ui = 12
	const uf = 3.14

	const (
		A = iota + 10 //0+10
		B
		C
	)
	fmt.Println(ui)
	fmt.Println(uf)
	fmt.Println(A, B, C)
}
