package find_the_index_of_the_first_occurrence_in_a_string

import (
	"strings"
)

func StrStr(haystack string, needle string) int {
	if len(haystack) == 1 && len(needle) == 1 {
		return 0
	}
	flag := strings.Contains(haystack, needle)

	mapK := make(map[int]int)
	if flag {
		ch := strings.Split(needle, "")
		for i := range haystack {
			if string(haystack[i]) == ch[0] {
				mapK[i] = i
			}
		}
		for i := 0; i < len(haystack); i++ {
			counter := i
			v, ex := mapK[i]
			if ex {
				flagv := true
				for j := range needle {
					if needle[j] != haystack[counter] {
						flagv = false
					}
					counter++
				}
				if flagv {
					return v
				}
			}

		}

	}

	return -1
}
