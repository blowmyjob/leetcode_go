package leetcode

import "sort"

func thirdMax(nums []int) int {
	sortNum := []int{}
	sortMap := make(map[int]int)
	sort.Ints(nums)
	for _, v := range nums {
		_, ok := sortMap[v]
		if !ok {
			sortMap[v] = 1
			sortNum = append(sortNum, v)
		}
	}
	if len(sortNum) < 3 {
		return sortNum[len(sortNum)-1]
	}
	return sortNum[len(sortNum)-3]
}
