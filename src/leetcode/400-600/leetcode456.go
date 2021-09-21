package leetcode

import "math"

func find132pattern(nums []int) bool {
	stack := []int{}
	min := math.MinInt64
	for i := len(nums) - 1; i >= 0; i-- {
		if min > nums[i] {
			return true
		}
		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			min = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
