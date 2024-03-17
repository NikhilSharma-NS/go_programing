package find_a_number

func Non_recursive_Find_A_Num(list []int, num int) bool {

	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
