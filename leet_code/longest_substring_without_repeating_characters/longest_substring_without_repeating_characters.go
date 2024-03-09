package longest_substring_without_repeating_characters

import "math"

func Longest_substring_without_repeating_characters(s string) int {

	mapK := make(map[byte]bool)
	first, second := 0, 0
	ans := 1.0

	for second < len(s) {
		c := s[second]
		second++
		for first < second {
			_, Ispresnt := mapK[c]
			if !Ispresnt {
				break
			}
			lowC := s[first]
			first++
			delete(mapK, lowC)

		}
		mapK[c] = true
		ans = math.Max(float64(ans), float64(second-first))

	}

	return int(ans)

}
