package main

import (
	"Gorepo/go_programing/learn_golang/alogo_with_go/find_a_number/non_recursive"
	"Gorepo/go_programing/learn_golang/alogo_with_go/reverse_a_string"
	"fmt"
)

func main() {

	isFound := non_recursive.Non_recursive_Find_A_Num([]int{1, 2, 3, 4, 5, 5}, 2)
	fmt.Println(isFound)

	rev_with_rune := reverse_a_string.Reverse_a_string_with_Rune("abcdefg")
	fmt.Println(rev_with_rune)

	reverse_a_string_with_byte := reverse_a_string.Reverse_a_string_with_byte("abcdefg")
	fmt.Println(reverse_a_string_with_byte)

	reverse_a_string_with_string_builder := reverse_a_string.Reverse_a_string_with_string_builder("abcdefg")
	fmt.Println(reverse_a_string_with_string_builder)
}
