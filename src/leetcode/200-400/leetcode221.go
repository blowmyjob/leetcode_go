package leetcode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix[0]) == 0 {
		return 0
	}
	var min func(i, j int) int
	var max func(i, j int) int
	min = func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}
	maxRes := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = int(matrix[i][j] - '0')
			}
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			maxRes = max(maxRes, dp[i][j])
		}
	}
	return maxRes * maxRes
}
