package leetcode

func increasingTriplet(nums []int) bool {
	n := 3
	if n > len(nums) {
		return false
	}
	res := make([]int, n-1)
	for i := range res {
		res[i] = 1<<31 - 1
	}
	for _, num := range nums {
		i := 0
		for ; i < n-1; i++ {
			if num < res[i] {
				res[i] = num
				break
			}
		}
		if i >= n-1 {
			return true
		}
	}
	return false
}
