package leetcode

func minCut(s string) int {
	n := len(s)
	isPali := make([][]bool, n)
	for i := range isPali {
		isPali[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				isPali[i][j] = true
			} else if j-i == 1 && s[i] == s[j] {
				isPali[i][j] = true
			} else if j-i > 1 && s[i] == s[j] && isPali[i+1][j-1] {
				isPali[i][j] = true
			}
		}
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	for i := 1; i < n; i++ {
		if isPali[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPali[j+1][i] {
				if dp[j]+1 < dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
	}
	return dp[n-1]
}
