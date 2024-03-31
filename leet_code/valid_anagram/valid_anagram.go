package valid_anagram

func IsAnagram(s string, t string) bool {
	flag := true
	mapv := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		v, ex := mapv[s[i]]
		if ex {
			mapv[s[i]] = v + 1
		} else {
			mapv[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		v, ex := mapv[t[i]]
		if ex {
			if v == 1 {
				delete(mapv, t[i])
			} else if v > 1 {
				mapv[t[i]] = v - 1
			}
		} else {
			flag = false
		}

	}

	return flag
}

func IsAnagramBestSolution(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var charCount [26]int

	for _, char := range s {
		charCount[char-'a']++
	}
	for _, char := range t {
		charCount[char-'a']--
		if charCount[char-'a'] < 0 {
			return false
		}
	}

	return true
}
