// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var arr []int
	arr = append([]int{64, 25, 12, 22}, 11)
	out := insertionsort(arr)
	fmt.Println(out)
}

func insertionsort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1

		for j := i && arr[j] > temp; j >= 0; j-- {
			arr[j+1] = arr[j]

		}
		fmt.Println(temp)
	}
	return arr

}
