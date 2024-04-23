package jump_game

func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	needs := 1
	for i := len(nums) - 2; i > 0; i-- {
		if nums[i] >= needs {
			needs = 1
			continue
		}
		needs++
	}

	return nums[0] >= needs
}
