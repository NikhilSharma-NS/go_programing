package quick_sort

func QuickSort(arr []int) []int {

	low := 0
	high := len(arr) - 1
	processQuickSort(arr, low, high)
	return arr

}

func processQuickSort(arr []int, low int, high int) {

	if low < high {
		pivot_index := partation(arr, low, high)

		processQuickSort(arr, low, pivot_index-1)
		processQuickSort(arr, pivot_index, high)
	}

}

func partation(arr []int, low int, high int) int {

	pivit_element := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivit_element {
			i++
			//swap
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp

		}
	}
	i++
	tem_piv := arr[i]
	arr[i] = pivit_element
	arr[high] = tem_piv
	return i
}
