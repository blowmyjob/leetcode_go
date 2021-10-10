package leetcode

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n)
	dp[0] = 1
	index := make([]int, len(primes))
	for i := 2; i <= n; i++ {
		min := math.MaxInt32
		for j := 0; j < len(primes); j++ {
			if min > primes[j]*dp[index[j]] {
				min = primes[j] * dp[index[j]]
			}
		}
		for j := 0; j < len(primes); j++ {
			if min == primes[j]*dp[index[j]] {
				index[j]++
			}
		}
		dp[i] = min
	}
	return dp[n-1]
}
