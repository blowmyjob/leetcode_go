package leetcode

func Rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}
	if n == 2 {
		return dpNum(nums[0], nums[1])
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = nums[0], dpNum(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = dpNum(dp[i-1], dp[i-2]+nums[i])
	}
	return dpNum(dp[n-1], dp[n-2])
}

func dpNum(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}
