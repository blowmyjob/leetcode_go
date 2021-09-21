package contest

import (
	"sort"
)

func EliminateMaximum(dist []int, speed []int) int {
	timeRate := make([]int, len(dist))
	for i := 0; i < len(dist); i++ {
		timeRate[i] = (dist[i] - 1) / speed[i]
	}
	sort.Ints(timeRate)
	for i := 0; i < len(dist); i++ {
		if timeRate[i] < i {
			return i
		}
	}
	return 0
}
