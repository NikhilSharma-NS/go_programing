package firstuniquecharacterinastring

func FirstUniqChar(s string) int {
	// This is a constant space allocation (ie: always 26)
	var counts = make([]int, 26)

	// Insert all the character's count into counts array
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}

	// Find the first element with count 1
	for i := 0; i < len(s); i++ {
		if counts[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func FirstUniqChar_Worst(s string) int {
	mapv := make(map[byte]int)
	final := -1
	for i := 0; i < len(s); i++ {
		key, isAvaliable := mapv[s[i]]
		if isAvaliable {
			mapv[s[i]] = key + 1
		} else {
			mapv[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {

		key, _ := mapv[s[i]]
		if key == 1 {
			final = i
			return final
		}
	}
	return final

}
