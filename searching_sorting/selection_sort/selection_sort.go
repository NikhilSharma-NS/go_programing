// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var arr []int
	arr = append([]int{64, 25, 12, 22}, 11)
	out := selectionsort(arr)
	fmt.Println(out)
}

func selectionsort(arr []int) []int {

	for counter := 0; counter < len(arr); counter++ {
		var min_index int
		min_index = counter
		for counter1 := counter + 1; counter1 < len(arr); counter1++ {
			if arr[min_index] > arr[counter1] {
				min_index = counter1
			}
		}
		temp := arr[min_index]
		arr[min_index] = arr[counter]
		arr[counter] = temp

	}
	return arr

}
