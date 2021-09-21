package leetcode

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		weight := right - left
		minHeight := minHeight(height[right], height[left])
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
		if weight*minHeight > max {
			max = weight * minHeight
		}
	}
	return max
}

func minHeight(i, j int) int {
	if i > j {
		return j
	}
	return i
}
