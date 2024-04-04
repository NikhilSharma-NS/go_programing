package majority_element

func MajorityElement(nums []int) int {

	m := make(map[int]int)
	mct := 0
	rs := nums[0]
	for _, v := range nums {
		ctr, exists := m[v]
		if exists {
			m[v] = ctr + 1
			if ctr > mct {
				mct = ctr
				rs = v
			}
		} else {
			m[v] = 1
		}

	}
	return rs

}

func MajorityElementBest(nums []int) int {
	count := 0
	res := 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if v == res {
			count++
		} else {
			count--
		}

	}
	return res

}
