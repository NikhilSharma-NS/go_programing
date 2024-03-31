package maximum_subarray

func MaxSubArray(nums []int) int {

	max := nums[0]
	if len(nums) == 1 {
		return max
	}
	currSum := 0
	for i := 0; i < len(nums); i++ {
		if currSum < 0 {
			currSum = 0
		}
		currSum = currSum + nums[i]
		if currSum > max {
			max = currSum
		}
	}
	return max

}
