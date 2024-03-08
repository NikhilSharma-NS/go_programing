package container_with_most_water

func MaxArea_solution1(height []int) int {
	var left, sum int
	for i := 0; i < len(height); i++ {
		var sum_v int
		for j := i + 1; j < len(height); j++ {
			if height[i] > height[j] {
				left = height[j]
			} else {
				left = height[i]
			}
			sum_v = (j - i) * left
			if sum_v > sum {
				sum = sum_v
			}
		}
	}

	return sum

}

func MaxArea_solution2(height []int) int {
	length := len(height)
	var sum int
	for i, j := 0, length-1; i < length && i != j; {
		var inner int
		if height[i] > height[j] {
			inner = (j - i) * height[j]
			j--
		} else {
			inner = (j - i) * height[i]
			i++
		}
		if sum < inner {
			sum = inner
		}

	}
	return sum
}
