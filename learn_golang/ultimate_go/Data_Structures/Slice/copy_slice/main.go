package main

import "fmt"

func main() {

	slice1 := []string{"A", "B", "C"}

	slice2 := make([]string, len(slice1))

	copy(slice2, slice1)

	inspacetion(slice1)
	inspacetion(slice2)

	// 	Len,Cap 3 3
	// Index[0]Add[0xc000110480]Value[A]
	// Index[1]Add[0xc000110490]Value[B]
	// Index[2]Add[0xc0001104a0]Value[C]
	// Len,Cap 3 3
	// Index[0]Add[0xc0001104b0]Value[A]
	// Index[1]Add[0xc0001104c0]Value[B]
	// Index[2]Add[0xc0001104d0]Value[C]
}

func inspacetion(slice []string) {
	fmt.Println("Len,Cap", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("Index[%d]Add[%p]Value[%s]\n", i, &slice[i], s)
	}
}
