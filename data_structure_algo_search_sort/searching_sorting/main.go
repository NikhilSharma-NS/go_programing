package main

import (
	"Gorepo/go_programing/data_structure_algo_search_sort/searching_sorting/merge_sort"
	"Gorepo/go_programing/data_structure_algo_search_sort/searching_sorting/quick_sort"
	"fmt"
)

func main() {

	quick_sorted_array := quick_sort.QuickSort([]int{6, 3, 9, 5, 2, 8})
	fmt.Println("Input Array  Using Quick Sort::[6, 3, 9, 5, 2, 8]")
	fmt.Println("Output Array Using Quick Sort:", quick_sorted_array)

	merged_sorted_array := merge_sort.MergeSort([]int{6, 3, 9, 5, 2, 8})
	fmt.Println("Input  Array Using Merge Sort::[6, 3, 9, 5, 2, 8]")
	fmt.Println("Output Array Using Merge Sort", merged_sorted_array)

}
