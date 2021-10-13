package leetcode

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	swag := 0
	maxlen := 2
	if nums[0] == nums[1] {
		swag = 0
		maxlen = 1
	}
	if nums[0] > nums[1] {
		swag = -1
	}
	if nums[0] < nums[1] {
		swag = 1
	}
	for i := 1; i < len(nums)-1; i++ {
		tmp := 0
		if nums[i] > nums[i+1] {
			tmp = -1
		} else if nums[i] < nums[i+1] {
			tmp = 1
		}
		if tmp == swag || tmp == 0 {
			continue
		} else {
			swag = tmp
			maxlen++
		}
	}
	return maxlen
}
