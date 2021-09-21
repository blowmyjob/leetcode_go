package leetcode

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := (left + right) / 2
		if nums[mid] > nums[left] {
			left = mid + 1
		} else if nums[mid] == nums[left] {
			left++
		} else {
			right--
		}
	}
	return nums[left]
}
