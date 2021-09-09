package leetcode

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	cur := intervals[0]
	end := intervals[0][1]
	for _, v := range intervals {
		if v[0] <= end {
			end = maxIndex(v[1], cur[1])
			cur = []int{cur[0], end}
		} else {
			res = append(res, cur)
			cur = v
			end = v[1]
		}
	}
	res = append(res, cur)
	return res
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}
