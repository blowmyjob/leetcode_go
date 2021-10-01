package leetcode

func minSubArrayLen(target int, nums []int) int {
	left := 0
	min := len(nums) + 1
	for left < len(nums) {
		sum := nums[left]
		right := left + 1
		count := 1
		if sum >= target {
			if count < min {
				min = count
				break
			}
		}
		for right < len(nums) {
			sum += nums[right]
			right++
			count++
			if sum >= target {
				if count < min {
					min = count
					break
				}
			}
		}
		left++
	}
	if min == len(nums)+1 {
		min = 0
	}
	return min
}
