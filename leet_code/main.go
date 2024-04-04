package main

import (
	"Gorepo/go_programing/leet_code/add_two_numbers"
	"Gorepo/go_programing/leet_code/best_time_to_buy_and_sell_stock"
	"Gorepo/go_programing/leet_code/check_if_array_is_sorted_and_rotated"
	"Gorepo/go_programing/leet_code/container_with_most_water"
	"Gorepo/go_programing/leet_code/contains_duplicate"
	"Gorepo/go_programing/leet_code/delete_node_in_a_linked_list"
	"Gorepo/go_programing/leet_code/delete_the_middle_node_of_a_linked_list"
	"Gorepo/go_programing/leet_code/first_unique_character_in_a_string"
	"Gorepo/go_programing/leet_code/generate_parentheses"
	"Gorepo/go_programing/leet_code/group_anagrams"
	"Gorepo/go_programing/leet_code/length_of_last_word"
	"Gorepo/go_programing/leet_code/longest_palindromic_substring"
	"Gorepo/go_programing/leet_code/longest_substring_without_repeating_characters"
	"Gorepo/go_programing/leet_code/maximum_subarray"
	"Gorepo/go_programing/leet_code/merge_intervals"
	"Gorepo/go_programing/leet_code/merge_sorted_array"
	"Gorepo/go_programing/leet_code/palindrome_number"
	"Gorepo/go_programing/leet_code/product_of_array_except_self"
	"Gorepo/go_programing/leet_code/remove_duplicates_from_sorted_array"
	"Gorepo/go_programing/leet_code/remove_duplicates_from_sorted_array_ii"
	"Gorepo/go_programing/leet_code/remove_element"
	"Gorepo/go_programing/leet_code/remove_linked_list_elements"
	"Gorepo/go_programing/leet_code/reverse_integer"
	"Gorepo/go_programing/leet_code/rotate_list"
	"Gorepo/go_programing/leet_code/threesum"
	"Gorepo/go_programing/leet_code/two_sum"
	"Gorepo/go_programing/leet_code/valid_anagram"
	"Gorepo/go_programing/leet_code/valid_parentheses"
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

	fmt.Println(first_unique_character_in_a_string.FirstUniqChar("aabbc"))

	fmt.Println(merge_sorted_array.Merge_sorted_array([]int{1, 2, 3}, 3, []int{2, 5, 6}, 3))

	fmt.Println(merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
	fmt.Println(merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{1}, 1, []int{}, 0))
	fmt.Println(merge_sorted_array.Merge_sorted_array_without_using_another_array([]int{0}, 0, []int{1}, 1))

	fmt.Println(remove_element.RemoveElementWithRange([]int{3, 2, 2, 3}, 3))
	fmt.Println(remove_element.RemoveElementWithRange([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Println(remove_element.RemoveElementWithRange_ReturnSlice([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

	fmt.Println(remove_duplicates_from_sorted_array.RemoveDuplicates_WithRange([]int{1, 1, 2}))

	fmt.Println(contains_duplicate.ContainsDuplicate([]int{2, 14, 18, 22, 22}))

	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{7, 6, 4, 3, 1}))

	fmt.Println(best_time_to_buy_and_sell_stock.MaxProfit([]int{1, 2}))

	fmt.Println(threesum.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))

	fmt.Println("Input: 123", reverse_integer.ReverseInteger(123))
	fmt.Println("Input: -123", reverse_integer.ReverseInteger(-123))
	fmt.Println("Input: 1534236469", reverse_integer.ReverseInteger(1534236469))

	fmt.Println("Input:[3,4,5,1,2]", check_if_array_is_sorted_and_rotated.Check_if_array_is_sorted_and_rotated([]int{3, 4, 5, 1, 2}))
	fmt.Println("Input:[2,1,3,4]", check_if_array_is_sorted_and_rotated.Check_if_array_is_sorted_and_rotated([]int{2, 1, 3, 4}))
	fmt.Println("Input:[1,2,3]", check_if_array_is_sorted_and_rotated.Check_if_array_is_sorted_and_rotated([]int{1, 2, 3}))

	head := &remove_linked_list_elements.ListNode{Val: 1, Next: &remove_linked_list_elements.ListNode{Val: 2}}
	fmt.Println(remove_linked_list_elements.RemoveElements(head, 1))

	node1 := &delete_node_in_a_linked_list.ListNode{Val: 4, Next: &delete_node_in_a_linked_list.ListNode{Val: 5, Next: &delete_node_in_a_linked_list.ListNode{Val: 1, Next: &delete_node_in_a_linked_list.ListNode{Val: 9}}}}
	fmt.Println("Before Deletion:")
	delete_node_in_a_linked_list.PrintData(node1)
	delete_node_in_a_linked_list.DeleteNode(node1.Next)
	fmt.Println()
	fmt.Println("After Deletion:")
	delete_node_in_a_linked_list.PrintData(node1)
	fmt.Println()

	node2 := &delete_the_middle_node_of_a_linked_list.ListNode{Val: 1, Next: &delete_the_middle_node_of_a_linked_list.ListNode{Val: 3, Next: &delete_the_middle_node_of_a_linked_list.ListNode{Val: 4}}}
	fmt.Println("Before Deletion of mid:")
	delete_the_middle_node_of_a_linked_list.PrintData(node2)
	delete_the_middle_node_of_a_linked_list.DeleteMiddle(node2)
	fmt.Println()
	fmt.Println("After Deletion of mid:")
	delete_the_middle_node_of_a_linked_list.PrintData(node2)
	fmt.Println()
	fmt.Println("Before Rotaion of [1,2,3,4,5]: by 2")
	node3 := &rotate_list.ListNode{Val: 1, Next: &rotate_list.ListNode{Val: 2, Next: &rotate_list.ListNode{Val: 3, Next: &rotate_list.ListNode{Val: 4, Next: &rotate_list.ListNode{Val: 5}}}}}
	node3 = rotate_list.RotateRight(node3, 2)
	rotate_list.PrintData(node3)

	fmt.Println("[aaa][bbb] is valid anagram or not: ", valid_anagram.IsAnagramBestSolution("aaa", "abb"))

	fmt.Println("valid-parentheses:()", valid_parentheses.IsValid("()"))
	fmt.Println("valid-parentheses:(]", valid_parentheses.IsValid("(]"))
	fmt.Println("valid-parentheses:{[]}:", valid_parentheses.IsValid("{[]}"))

	fmt.Println("maximum subarray of [-2, 1, -3, 4, -1, 2, 1, -5, 4]: ", maximum_subarray.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	fmt.Println("maximum subarray of [1] : ", maximum_subarray.MaxSubArray([]int{1}))

	fmt.Println("product_of_array_except_self [1,2,3,4]::", product_of_array_except_self.ProductExceptSelf([]int{1, 2, 3, 4}))

	fmt.Println("len of last word", length_of_last_word.LengthOfLastWord_Best("   fly me   to   the moon  "))
	fmt.Println("len of last word", length_of_last_word.LengthOfLastWord_Best("Hello World"))
	fmt.Println("len of last word", length_of_last_word.LengthOfLastWord_Best("a"))

	vales := [][]int{{1, 3}, {8, 10}, {15, 18}, {2, 6}}

	fmt.Println("merge_intervals{1, 3}, {8, 10}, {15, 18}, {2, 6} to ", merge_intervals.Merge(vales))

	s := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	v := group_anagrams.GroupAnagrams(s)
	fmt.Println("Group of anagram", v)

	gen3 := generate_parentheses.GenerateParenthesis(3)
	fmt.Println("GenerateParenthesis", gen3)

	gen1 := generate_parentheses.GenerateParenthesis(1)
	fmt.Println("GenerateParenthesis", gen1)

	arr_removeDuplicates := []int{1, 1, 2, 2, 2, 3, 6, 7, 8, 9, 9, 9, 9, 9, 9}

	left := remove_duplicates_from_sorted_array_ii.RemoveDuplicates(arr_removeDuplicates)
	fmt.Println()
	for i := 0; i < left; i++ {
		fmt.Printf("[%v]", arr_removeDuplicates[i])
	}

	two_sum := two_sum.TwoSum([]int{1, 2, 3, 4, 5, 6}, 10)

	fmt.Println("Two Sum of [1, 2, 3, 4, 5, 6],10", two_sum)

	listNode1 := add_two_numbers.ListNode{Val: 1, Next: &add_two_numbers.ListNode{Val: 2}}
	listNode2 := add_two_numbers.ListNode{Val: 2, Next: &add_two_numbers.ListNode{Val: 4}}

	li := add_two_numbers.AddTwoNumbers(&listNode1, &listNode2)
	fmt.Println("Sum link list")
	li.PrintData()

}
