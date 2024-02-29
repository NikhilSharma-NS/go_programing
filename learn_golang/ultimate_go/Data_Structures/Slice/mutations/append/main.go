package main

import "fmt"

func main() {

	slice := []string{"A", "B", "C", "D", "E"}

	slice2 := slice[2:4]

	slice2 = append(slice2, "Change")

	inspacetion(slice)
	inspacetion(slice2)

	//Index[0]Add[0xc0000d6000]Value[A]
	//Index[1]Add[0xc0000d6010]Value[B]
	//Index[2]Add[0xc0000d6020]Value[C]
	//Index[3]Add[0xc0000d6030]Value[D]
	//Index[4]Add[0xc0000d6040]Value[Change]
	//Len,Cap 3 3
	//Index[0]Add[0xc0000d6020]Value[C]
	//Index[1]Add[0xc0000d6030]Value[D]
	//Index[2]Add[0xc0000d6040]Value[Change]

	slice3 := []string{"A", "B", "C", "D", "E"}

	slice4 := slice3[2:4:4]
	inspacetion(slice3)
	inspacetion(slice4)
	//[a:b:c]
	// len(a-b)
	//cap(a-c)

	// 	Len,Cap 5 5
	// Index[0]Add[0xc00006a0a0]Value[A]
	// Index[1]Add[0xc00006a0b0]Value[B]
	// Index[2]Add[0xc00006a0c0]Value[C]
	// Index[3]Add[0xc00006a0d0]Value[D]
	// Index[4]Add[0xc00006a0e0]Value[E]
	// Len,Cap 2 2
	// Index[0]Add[0xc00006a0c0]Value[C]
	// Index[1]Add[0xc00006a0d0]Value[D]

}

func inspacetion(slice []string) {
	fmt.Println("Len,Cap", len(slice), cap(slice))

	for i, s := range slice {
		fmt.Printf("Index[%d]Add[%p]Value[%s]\n", i, &slice[i], s)
	}
}
