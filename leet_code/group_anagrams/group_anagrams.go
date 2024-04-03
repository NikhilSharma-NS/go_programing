package group_anagrams

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func GroupAnagrams(strs []string) (output [][]string) {
	mv := make(map[string][]string)
	for _, v := range strs {
		x := SortString(v)
		key, ex := mv[x]
		if ex {
			key = append(key, v)
			mv[x] = key
		} else {
			mv[x] = []string{v}
		}
	}

	for _, value := range mv {
		output = append(output, value)
	}

	return output
}

func word_stats(w string) [26]uint16 {
	res := [26]uint16{}
	for _, c := range w {
		res[c-'a']++
	}

	return res
}

func GroupAnagramsBest(words []string) [][]string {
	data := map[[26]uint16][]string{}

	for _, w := range words {
		stats := word_stats(w)
		data[stats] = append(data[stats], w)
	}

	res, p := make([][]string, len(data)), 0
	for _, val := range data {
		res[p] = val
		p++
	}
	return res
}
