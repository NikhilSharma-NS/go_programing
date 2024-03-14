package main

import "fmt"

type Numbers []int

type operatefun[T any] func(t T) T

func opearte[T any](slice []T, fn operatefun[T]) []T {
	re := make([]T, len(slice))

	for i, v := range slice {
		re[i] = fn(v)
	}
	return re

}

type Slice[T any] interface {
	type []T
}

func opearte1[S Slice[T]](slice S, fn operatefun[T]) S {
	re := make(S, len(slice))

	for i, v := range slice {
		re[i] = fn(v)
	}
	return re

}

func Double(n Numbers) Numbers {

	fn := func(n int) int {
		return 2 * n
	}

	ns := opearte(n, fn)
	fmt.Printf("%T", ns)
	fmt.Printf("%v", ns)
	return ns
}
func main() {
	x := []int{1, 2}
	Double(x)
}
