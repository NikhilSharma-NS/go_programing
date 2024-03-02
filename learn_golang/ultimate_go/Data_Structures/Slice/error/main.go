package main

import "fmt"

func main() {
	friends := []string{"A", "B", "C", "D", "E", "F"}

	for _, v := range friends {
		friends = friends[:2]
		fmt.Println("vavlue", v)
	}
	friends1 := []string{"A", "B", "C", "D", "E", "F", "G"}
	for i := range friends1 {
		fmt.Println("Index", i)
		friends1 = friends1[:2]
		fmt.Println("vavlue", friends1[i])
	}

	panic: runtime error: index out of range [2] with length 2

// goroutine 1 [running]:
// main.main()
//         **/Data_Structures/Slice/error/main.go:16 +0x389
// exit status 2
}
