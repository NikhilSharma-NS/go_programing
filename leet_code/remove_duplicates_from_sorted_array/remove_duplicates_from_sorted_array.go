package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

func RemoveDuplicates_WithRange(nums []int) int {
	l := 0
	for _, v := range nums {
		if nums[l] != v {
			l++
			nums[l] = v
		}
	}
	return l + 1

}
