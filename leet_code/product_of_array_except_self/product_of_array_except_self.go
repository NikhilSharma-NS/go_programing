package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	pre := 1
	post := 1

	for i := 0; i < len(nums); i++ {
		res[i] = pre
		pre = nums[i] * pre
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * post
		post = nums[i] * post
	}

	return res
}
