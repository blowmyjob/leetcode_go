package leetcode

func findDuplicate(nums []int) int {
	x := 0
	for nums[x] != 0 {
		x, nums[x] = nums[x], 0
	}
	return x
}
