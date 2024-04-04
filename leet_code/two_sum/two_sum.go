package two_sum

func TwoSum(nums []int, target int) []int {

	var output []int
	for counter := 0; counter < len(nums); counter++ {
		for counter1 := counter + 1; counter1 < len(nums); counter1++ {
			if nums[counter]+nums[counter1] == target {
				output = append(output, counter)
				output = append(output, counter1)
				return output
				break
			}
		}
	}
	return output

}
