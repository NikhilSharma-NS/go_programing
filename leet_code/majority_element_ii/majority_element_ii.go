package majority_element_ii

func MajorityElements(nums []int) (res []int) {
	m := make(map[int]int)
	for _, v := range nums {
		ctr, exists := m[v]
		if exists {
			m[v] = ctr + 1
		} else {
			m[v] = 1
		}
	}

	for i, v := range m {
		if v > len(nums)/3 {
			res = append(res, i)
		}
	}

	return res
}
func MajorityElementsBest(nums []int) (res []int) {
	// We can have only 2 numbers with count greater than len(nums)/3
	// Find those 2 possible numbers first
	c1, c2 := 0, 0
	mej1, mej2 := 0, 0
	for _, v := range nums {
		if v == mej1 {
			c1++
		} else if v == mej2 {
			c2++
		} else if c1 == 0 {
			mej1, c1 = v, 1
		} else if c2 == 0 {
			mej2, c2 = v, 1
		} else {
			c1--
			c2--
		}

	}

	// Verify count of those 2 possible numbers again
	c1, c2 = 0, 0
	for _, num := range nums {
		if num == mej1 {
			c1++
		} else if num == mej2 {
			c2++
		}
	}

	// Add those numbers if they satisfy conditions
	if c1 > len(nums)/3 {
		res = append(res, mej1)
	}
	if c2 > len(nums)/3 {
		res = append(res, mej2)
	}
	return res
}
