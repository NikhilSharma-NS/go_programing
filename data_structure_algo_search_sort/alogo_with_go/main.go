package main

import (
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/any_base_to_decimal"
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/convert_decimal_to_any_base"
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/find_a_number"
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/find_two_num_in_list_sum"
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/fizz_buzz"
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/reverse_a_string"
	"Gorepo/go_programing/data_structure_algo_search_sort/alogo_with_go/sum_of_number"
	"fmt"
)

func main() {

	isFound := find_a_number.Non_recursive_Find_A_Num([]int{1, 2, 3, 4, 5, 5}, 2)
	fmt.Println(isFound)

	rev_with_rune := reverse_a_string.Reverse_a_string_with_Rune("abcdefg诶")
	fmt.Println(rev_with_rune)

	reverse_a_string_with_byte := reverse_a_string.Reverse_a_string_with_byte("abcdefg")
	fmt.Println(reverse_a_string_with_byte)

	reverse_a_string_with_string_builder := reverse_a_string.Reverse_a_string_with_string_builder("abcdefg")
	fmt.Println(reverse_a_string_with_string_builder)

	sum_of_number_with_out_recurrsion := sum_of_number.Sum_of_number_with_out_recurrsion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(sum_of_number_with_out_recurrsion)

	sum_of_number_with_recurrsion := sum_of_number.Sum_of_number_with_recurrsion([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(sum_of_number_with_recurrsion)

	fizz_buzz.Fizz_buzz(20)
	fizz_buzz.Fizz_buzz_with_switch(20)
	fmt.Println("Base: 2", convert_decimal_to_any_base.ConvertDecimalToBase(2044, 2))
	fmt.Println("Base: 8", convert_decimal_to_any_base.ConvertDecimalToBase(2044, 8))
	fmt.Println("Base: 16", convert_decimal_to_any_base.ConvertDecimalToBase(2044, 16))
	fmt.Println("Base: 16", convert_decimal_to_any_base.ConvertDecimalToBase_Constant(2044, 16))
	fmt.Println("Base: 16", convert_decimal_to_any_base.ConvertDecimalToBase_Constant_with_String_Builder(2044, 16))

	fmt.Println("Base: 10", any_base_to_decimal.BasetoDecimal("11111111100", 2))
	fmt.Println("Base: 10", any_base_to_decimal.BasetoDecimal("7FC", 16))

	index1_worst, index2_worst := find_two_num_in_list_sum.Find_two_num_in_list_sum_worst([]int{1, 2, 3, 4}, 7)
	fmt.Printf("Index1: [%d], Index2: [%d]", index1_worst, index2_worst)
	fmt.Println()
	index1, index2 := find_two_num_in_list_sum.Find_two_num_in_list_sum([]int{1, 2, 3, 4}, 7)
	fmt.Printf("Index1: [%d], Index2: [%d]", index1, index2)

}
