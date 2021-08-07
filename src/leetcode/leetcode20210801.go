package leetcode

import "sort"

func numberOfWeeks(milestones []int) int64 {
	if len(milestones) == 1 {
		return 1
	}
	sort.Ints(milestones)
	res := 0
	for i := 0; i < len(milestones)-1; i++ {
		res += milestones[i]
	}
	if res < milestones[len(milestones)-1] {
		res *= 2
		res += 1
		return int64(res)
	}
	res *= 2
	return int64(res)
}
