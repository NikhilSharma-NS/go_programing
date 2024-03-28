package check_if_array_is_sorted_and_rotated

func Check_if_array_is_sorted_and_rotated(nums []int) bool {

	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count = count + 1
		}
	}
	if nums[len(nums)-1] > nums[0] {
		count++
	}
	return count <= 1

}
