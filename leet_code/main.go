package main

import (
	"Gorepo/go_programing/leet_code/container_with_most_water"
	firstuniquecharacterinastring "Gorepo/go_programing/leet_code/first_unique_character_in_a_string"
	"Gorepo/go_programing/leet_code/longest_palindromic_substring"
	"Gorepo/go_programing/leet_code/longest_substring_without_repeating_characters"
	"Gorepo/go_programing/leet_code/merge_sorted_array"
	"Gorepo/go_programing/leet_code/palindrome_number"
	"Gorepo/go_programing/leet_code/remove_element"
	"fmt"
)

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(container_with_most_water.MaxArea_solution1(arr))
	fmt.Println(container_with_most_water.MaxArea_solution2(arr))

	fmt.Println("longest_palindromic_substring", longest_palindromic_substring.LongestPalindrome("babad"))

	fmt.Println("Number is Palindrome : ", palindrome_number.Palindrome_number(1212))

	fmt.Println(longest_substring_without_repeating_characters.Longest_substring_without_repeating_characters("abcabcbb"))

	fmt.Println(longest_substring_without_repeating_characters.Longest_substring_without_repeating_characters("bbbbbbb"))

	fmt.Println(firstuniquecharacterinastring.FirstUniqChar("aabbc"))

	fmt.Println(merge_sorted_array.Merge_sorted_array([]int{1, 2, 3}, 3, []int{2, 5, 6}, 3))

	fmt.Println(merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	fmt.Println(merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{1}, 1, []int{}, 0))
	fmt.Println(merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{0}, 0, []int{1}, 1))

	fmt.Println(remove_element.RemoveElementWithRange([]int{3, 2, 2, 3}, 3))
	fmt.Println(remove_element.RemoveElementWithRange([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Println(remove_element.RemoveElementWithRange_ReturnSlice([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}
