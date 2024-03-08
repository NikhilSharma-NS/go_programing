package longest_palindromic_substring

import (
	"math"
)

func LongestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandFromMiddle(s, i, i)
		len2 := expandFromMiddle(s, i, i+1)
		lens := math.Max(float64(len1), float64(len2))
		if int(lens) > end-start {
			start = i - ((int(lens) - 1) / 2)
			end = i + (int(lens) / 2)
		}

	}
	return s[start : end+1]

}

func expandFromMiddle(s string, left int, right int) int {

	if len(s) == 0 || left > right {
		return 0
	}
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	resu := right - left - 1
	return resu

}
