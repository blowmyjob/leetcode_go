package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	diff := math.MaxInt32
	cloest := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			newDiff := int(math.Abs(float64(nums[i] + nums[left] + nums[right] - target)))
			if newDiff < diff {
				diff = newDiff
				cloest = sum
			}
			if sum > target {
				right--
			} else if sum <= target {
				left++
			}
		}
	}
	return cloest
}
