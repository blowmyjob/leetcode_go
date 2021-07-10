package leetcode

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] <= intervals[j][1]
	})
	res := 0
	rightEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= rightEnd {
			rightEnd = intervals[i][1]
		} else {
			res++
		}
	}
	return res
}
