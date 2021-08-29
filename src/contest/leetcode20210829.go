package contest

import (
	"math"
	"sort"
	"strings"
)

func KthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) == len(nums[j]) {
			return strings.Compare(nums[i], nums[j]) > 0
		}
		return len(nums[i]) > len(nums[j])
	})
	return nums[k-1]
}

func MinimumDifference(nums []int, k int) int {
	min := math.MaxInt32
	sort.Ints(nums)
	for i := 0; i <= len(nums)-k; i++ {
		minNum := nums[i]
		maxNum := nums[i+k-1]
		tmpNum := maxNum - minNum
		if tmpNum < min {
			min = tmpNum
		}
	}
	return min
}
