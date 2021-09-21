package leetcode

import "sort"

func findUnsortedSubarray(nums []int) int {
	snums := []int{}
	for _, v := range nums {
		snums = append(snums, v)
	}
	sort.Ints(snums)
	start, end := len(snums)-1, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != snums[i] {
			start = minIndex(start, i)
			end = maxIndex(end, i)
		}
	}
	if end > start {
		return end - start + 1
	}
	return 0
}

func minIndex(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxIndex(x, y int) int {
	if x < y {
		return y
	}
	return x
}
