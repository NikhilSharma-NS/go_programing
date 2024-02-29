package main

import "fmt"

func main() {

	var slice1 []string

	slice2 := []string{}

	slice3 := make([]string, 2)

	slice4 := make([]string, 2, 3)

	slice5 := []string{"A"}

	fmt.Println(slice1, slice2, slice3, slice4, slice5)
	fmt.Println(len(slice1), len(slice2), len(slice3), len(slice4), len(slice5))
	fmt.Println(cap(slice1), cap(slice2), cap(slice3), cap(slice4), cap(slice5))

}
