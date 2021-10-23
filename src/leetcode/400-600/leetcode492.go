package leetcode

import "math"

func constructRectangle(area int) []int {
	sqrt := math.Sqrt(float64(area))
	for i := int(sqrt); i >= 1; i-- {
		if area%i == 0 {
			return []int{i, area / i}
		}
	}
	return []int{}
}
