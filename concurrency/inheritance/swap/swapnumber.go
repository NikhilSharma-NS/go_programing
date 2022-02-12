package main

import "fmt"

func swapContents(listObj []int) {
	for i, j := 0, len(listObj)-1; i < j; i, j = i+1, j-1 {
		listObj[i], listObj[j] = listObj[j], listObj[i]
	}
}
func main() {
	listObj := []int{1, 2, 3, 7, 8}
	lis := []string{"a", "b", "c", "d"}
	swapContents(listObj)
	my(lis)
	fmt.Println(lis)
	fmt.Println(listObj)
}

func my(list []string) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}

}
