package leetcode

import "sort"

func firstMissingPositive(nums []int) int {
	wordMap := make(map[int]int)
	sort.Ints(nums)
	for _, v := range nums {
		wordMap[v] = v
	}
	for i := 1; i < len(nums); i++ {
		if _, ok := wordMap[i]; !ok {
			return i
		}
	}
	return len(nums) + 1
}
