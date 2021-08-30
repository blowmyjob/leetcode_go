package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := [][]int{}
	if n < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
			if sum > 0 {
				right--
			}
			if sum < 0 {
				left++
			}
		}
	}
	return res
}
