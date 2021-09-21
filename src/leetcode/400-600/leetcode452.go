package leetcode

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	ans := 1
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	maxRight := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > maxRight {
			maxRight = points[i][1]
			ans++
		}
	}
	return ans
}
