package main

import (
	"errors"
	"fmt"
)

type Vector[T any] []T

func (v Vector[T]) last() (T, error) {
	var zero T
	if len(v) == 0 {
		return zero, errors.New("empty")
	}
	return v[len(v)-1], nil
}

func main() {

	fmt.Println("vector [int] : ")
	t := []any{"2"}
	vGenInt := []Vector{t}

	//i, err := vGenInt.last()

	//fmt.Println(i, err)

}
