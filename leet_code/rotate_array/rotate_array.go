package rotate_array

import "fmt"

func Rotate(nums []int, k int) {
	if len(nums) == 1 || len(nums) == 0 {
		return
	}
	numsa := []int{}
	if k > len(nums) {
		k = k % len(nums)
	}
	sum := len(nums) - k

	for i := sum; i < len(nums); i++ {
		numsa = append(numsa, nums[i])
	}

	for j := 0; j < sum; j++ {
		numsa = append(numsa, nums[j])
	}

	for i := range numsa {
		nums[i] = numsa[i]
	}

	fmt.Println(nums)

}
