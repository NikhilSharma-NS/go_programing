package sum_of_number

func Sum_of_number_with_out_recurrsion(numbers []int) int {

	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func Sum_of_number_with_recurrsion(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum_of_number_with_out_recurrsion(numbers[1:])
}
