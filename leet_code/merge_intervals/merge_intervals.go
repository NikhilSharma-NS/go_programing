package merge_intervals

import (
	"fmt"
	"sort"
)

func Merge(intervals [][]int) [][]int {
	fmt.Println(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	n := len(intervals)
	currentStart := intervals[0][0]
	maxEnd := intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] > maxEnd {
			res = append(res, []int{currentStart, maxEnd})
			currentStart = intervals[i][0]
			maxEnd = intervals[i][1]
		} else {
			maxEnd = max(maxEnd, intervals[i][1])
		}
	}
	return append(res, []int{currentStart, maxEnd})
}
