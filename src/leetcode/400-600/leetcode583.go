package leetcode

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	}
	if len(word1) == 0 && len(word2) != 0 {
		return len(word2)
	}
	if len(word1) != 0 && len(word2) == 0 {
		return len(word1)
	}

	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				if dp[i+1][j] > dp[i][j+1] {
					dp[i+1][j+1] = dp[i+1][j]
				} else {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}
	return len(word1) + len(word2) - 2*dp[len(word1)][len(word2)]
}
