// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	var arr []int
	arr = append([]int{64, 25, 12, 22}, 11)
	out := bubblesort(arr)
	fmt.Println(out)
}

func bubblesort(arr []int) []int {
	var swapped bool
	for i := 0; i < len(arr); i++ {
		swapped = false
		for j := 0; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {

				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				swapped = true
			}

		}
		if !swapped {
			return arr
		}

	}
	return arr

}
