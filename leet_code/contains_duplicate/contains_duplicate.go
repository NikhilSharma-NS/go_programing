package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, value := range nums {
		i, is_Present := m[value]
		if !is_Present {
			m[value] = i + 1
		} else {
			return true
		}

	}

	return false
}
