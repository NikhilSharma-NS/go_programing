package main

import "fmt"

func main() {

	slice := []string{"A", "B", "C", "D", "E"}

	fmt.Println("Len,Cap", len(slice), cap(slice)) // 5 5

	inspacetion(slice)

	//Len,Cap 5 5
	//Len,Cap 5 5
	//Index[0]Add[0xc000100050]Value[A]
	//Index[1]Add[0xc000100060]Value[B]
	//Index[2]Add[0xc000100070]Value[C]
	//Index[3]Add[0xc000100080]Value[D]
	//Index[4]Add[0xc000100090]Value[E]

	slice2 := slice[2:4]

	fmt.Println("Len,Cap", len(slice2), cap(slice2)) // 2 3
	inspacetion(slice2)

	//Len,Cap 2 3
	//Len,Cap 2 3
	//Index[0]Add[0xc000100070]Value[C]
	//Index[1]Add[0xc000100080]Value[D]

}

func inspacetion(slice []string) {
	fmt.Println("Len,Cap", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("Index[%d]Add[%p]Value[%s]\n", i, &slice[i], s)
	}
}
