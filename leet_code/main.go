package main

import (
	"Gorepo/go_programing/leet_code/container_with_most_water"
	"Gorepo/go_programing/leet_code/longest_palindromic_substring"
	"Gorepo/go_programing/leet_code/palindrome_number"
	"fmt"
)

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(container_with_most_water.MaxArea_solution1(arr))
	fmt.Println(container_with_most_water.MaxArea_solution2(arr))

	s := "babad"
	fmt.Println(longest_palindromic_substring.LongestPalindrome(s))
	fmt.Println("Number is Palindrome : ", palindrome_number.Palindrome_number(1212))

}
