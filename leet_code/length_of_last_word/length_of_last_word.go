package length_of_last_word

import (
	"fmt"
	"strings"
)

func LengthOfLastWord(s string) int {

	chars := strings.Split(s, " ")
	count := 0

	for i := len(chars) - 1; i >= 0; i-- {
		fmt.Println("v", chars[i])
		if len(chars[i]) > 0 {
			count = len(chars[i])
			return count
		}
	}
	return count
}

func LengthOfLastWord_Best(s string) int {
	words := strings.Fields(strings.TrimSpace(s))
	return len(words[len(words)-1])
}
