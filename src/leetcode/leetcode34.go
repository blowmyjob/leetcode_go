package leetcode

func SearchRange(nums []int, target int) []int {
	find := midFind(nums, target)
	if find == -1 {
		return []int{-1, -1}
	}
	left, right := find-1, find+1
	for left >= 0 && nums[left] == target {
		left--
	}
	for right < len(nums) && nums[right] == target {
		right++
	}
	return []int{left + 1, right - 1}
}

func midFind(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
