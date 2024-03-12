package main

import "fmt"

func show[T any](slice []T) {
	fmt.Println("Generics")
	for _, v := range slice {
		fmt.Println(v, "")
	}
}

func main() {

	show([]string{"A", "B"})
	show([]int{1, 2})
}
