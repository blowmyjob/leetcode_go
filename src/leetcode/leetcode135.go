package leetcode

func candy(ratings []int) int {
	sum := 0
	left := []int{}
	for i := 0; i < len(ratings); i++ {
		left = append(left, 1)
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] && left[i-1] >= left[i] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && left[i+1] >= left[i] {
			left[i] = left[i+1] + 1
		}
	}
	for i := 0; i < len(left); i++ {
		sum += left[i]
	}
	return sum
}
