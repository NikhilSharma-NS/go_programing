package main

import "fmt"

func main() {

	slice := []string{"A", "B", "C", "D", "E"}

	slice2 := slice[2:4]

	slice2[0] = "Change"

	inspacetion(slice)
	inspacetion(slice2)

	//Len,Cap 5 5
	//Index[0]Add[0xc0000d6000]Value[A]
	//Index[1]Add[0xc0000d6010]Value[B]
	//Index[2]Add[0xc0000d6020]Value[Change]
	//Index[3]Add[0xc0000d6030]Value[D]
	//Index[4]Add[0xc0000d6040]Value[E]
	//Len,Cap 2 3
	//Index[0]Add[0xc0000d6020]Value[Change]
	//Index[1]Add[0xc0000d6030]Value[D]

}

func inspacetion(slice []string) {
	fmt.Println("Len,Cap", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("Index[%d]Add[%p]Value[%s]\n", i, &slice[i], s)
	}
}
