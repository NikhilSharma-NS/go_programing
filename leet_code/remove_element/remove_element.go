package remove_element

func RemoveElement(nums []int, val int) int {
	counter := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[counter] = nums[i]
			counter++
		}
	}

	return counter
}

func RemoveElementWithRange(nums []int, val int) int {
	v := 0

	for i, _ := range nums {
		if nums[i] != val {
			nums[v] = nums[i]
			v++
		}
	}

	return v

}

func RemoveElementWithRange_ReturnSlice(nums []int, val int) []int {
	v := 0

	for i, _ := range nums {
		if nums[i] != val {
			nums[v] = nums[i]
			v++
		}
	}
	return nums[:v]

}
