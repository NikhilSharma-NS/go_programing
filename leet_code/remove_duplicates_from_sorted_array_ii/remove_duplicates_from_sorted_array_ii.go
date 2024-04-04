package remove_duplicates_from_sorted_array_ii

import "math"

func RemoveDuplicates(nums []int) int {

	l := 0
	r := 0

	for r < len(nums) {
		count := 1

		for r+1 < len(nums) && nums[r] == nums[r+1] {
			count++
			r = r + 1
		}
		min := math.Min(float64(count), 2)
		for min > 0 {
			nums[l] = nums[r]
			l = l + 1
			min--

		}
		r = r + 1
	}
	return l

}
