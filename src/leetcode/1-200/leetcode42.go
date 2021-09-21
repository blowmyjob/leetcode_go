package leetcode

func trap(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if maxLeft < height[left] {
				maxLeft = height[left]
			}
			res += maxLeft - height[left]
			left++
		} else {
			if maxRight < height[right] {
				maxRight = height[right]
			}
			res += maxRight - height[right]
			right++
		}
	}
	return res
}
