package offer

func Fibonacci(n int) int {
	// write code here
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 3
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 3
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func jumpFloor(number int) int {
	// write code here
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 3
	}
	dp := make([]int, number+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 3
	for i := 2; i <= number; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[number]
}
