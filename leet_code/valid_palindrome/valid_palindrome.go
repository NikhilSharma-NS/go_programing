package valid_palindrome

import (
	"slices"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	slicess := strings.Split(s, "")
	charcters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i, j := 0, len(slicess)-1; i < len(slicess) && j >= 0; {
		if !slices.Contains(charcters, slicess[i]) {
			i = i + 1
			continue
		}
		if !slices.Contains(charcters, slicess[j]) {
			j = j - 1
			continue
		}

		if !strings.EqualFold(slicess[i], slicess[j]) {
			return false
		}
		i = i + 1
		j = j - 1

	}
	return true

}

func IsPalindromeOptimal(s string) bool {
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i = i + 1
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j = j - 1
			continue
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i = i + 1
		j = j - 1

	}
	return true
}
