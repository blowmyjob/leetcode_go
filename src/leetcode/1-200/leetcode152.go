package leetcode

import "math"

func maxProduct(nums []int) int {
	res := math.MinInt32
	min, max := 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		min = dfsMinProduct(nums[i], min*nums[i])
		max = dfsMaxProduct(nums[i], max*nums[i])
		res = dfsMaxProduct(res, max)
	}
	return res
}

func dfsMaxProduct(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func dfsMinProduct(i, j int) int {
	if i < j {
		return i
	}
	return j
}
