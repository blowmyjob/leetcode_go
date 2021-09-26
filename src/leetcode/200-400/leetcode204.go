package leetcode

func countPrimes(n int) int {
	count := 0
	dp := make([]bool, n+1)
	for i := 2; i < n; i++ {
		if dp[i] == false {
			count++
			for j := i + i; j < n; j += i {
				dp[j] = true
			}
		}
	}
	return count
}
