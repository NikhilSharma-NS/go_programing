package find_two_num_in_list_sum

func Find_two_num_in_list_sum(values []int, sum int) (index1 int, index2 int) {
	mapTemp := make(map[byte]int)
	for i := 0; i < len(values); i++ {
		v, ispresent := mapTemp[byte(sum-values[i])]
		if ispresent {
			return v, i
		} else {
			mapTemp[byte(values[i])] = i
		}
	}
	return -1, -1
}

func Find_two_num_in_list_sum_worst(values []int, sum int) (index1 int, index2 int) {

	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i]+values[j] == sum {
				return i, j
			}
		}
	}
	return -1, -1
}
