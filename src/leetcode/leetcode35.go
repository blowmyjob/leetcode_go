package leetcode

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			ans = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return ans
}
