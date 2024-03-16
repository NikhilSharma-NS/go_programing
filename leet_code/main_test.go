package main

import (
	"Gorepo/go_programing/leet_code/longest_palindromic_substring"
	"Gorepo/go_programing/leet_code/merge_sorted_array"
	"testing"
)

func TestLongest_palindromic_substring(t *testing.T) {
	longest_palindromic_substring.LongestPalindrome("babad")
	merge_sorted_array.Merge_sorted_array([]int{1, 2, 3}, 3, []int{2, 5, 6}, 3)
	merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}
