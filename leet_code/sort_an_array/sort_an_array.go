package sort_an_array

func SortArrayWorst(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}

func SortArrayBest(nums []int) []int {
	divide(nums, 0, len(nums)-1)
	return nums

}

func divide(arr []int, start_index int, end_index int) {
	if start_index >= end_index {
		return
	}
	mid_index := start_index + (end_index-start_index)/2

	divide(arr, start_index, mid_index)
	divide(arr, mid_index+1, end_index)
	conquer(arr, start_index, mid_index, end_index)

}

func conquer(arr []int, start_index int, mid_index int, end_index int) {

	merg_arr := []int{}
	index1 := start_index
	index2 := mid_index + 1

	for index1 <= mid_index && index2 <= end_index {
		if arr[index1] <= arr[index2] {
			merg_arr = append(merg_arr, arr[index1])
			index1++
		} else {
			merg_arr = append(merg_arr, arr[index2])
			index2++

		}
	}

	for index1 <= mid_index {
		merg_arr = append(merg_arr, arr[index1])
		index1++
	}
	for index2 <= end_index {
		merg_arr = append(merg_arr, arr[index2])
		index2++
	}

	for i, j := 0, start_index; i < len(merg_arr); i, j = i+1, j+1 {
		arr[j] = merg_arr[i]
	}

}
