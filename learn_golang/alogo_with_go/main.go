package main

import (
	"Gorepo/go_programing/learn_golang/alogo_with_go/find_a_number"
	"Gorepo/go_programing/learn_golang/alogo_with_go/reverse_a_string"
	"Gorepo/go_programing/learn_golang/alogo_with_go/sum_of_number"
	"fmt"
)

func main() {

	isFound := find_a_number.Non_recursive_Find_A_Num([]int{1, 2, 3, 4, 5, 5}, 2)
	fmt.Println(isFound)

	rev_with_rune := reverse_a_string.Reverse_a_string_with_Rune("abcdefg")
	fmt.Println(rev_with_rune)

	reverse_a_string_with_byte := reverse_a_string.Reverse_a_string_with_byte("abcdefg")
	fmt.Println(reverse_a_string_with_byte)

	reverse_a_string_with_string_builder := reverse_a_string.Reverse_a_string_with_string_builder("abcdefg")
	fmt.Println(reverse_a_string_with_string_builder)

	sum_of_number_with_out_recurrsion := sum_of_number.Sum_of_number_with_out_recurrsion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(sum_of_number_with_out_recurrsion)

	sum_of_number_with_recurrsion := sum_of_number.Sum_of_number_with_recurrsion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(sum_of_number_with_recurrsion)
}
