package leetcode

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 1; i <= amount; i++ {
		dp[i] = 999999
		for j := 0; j < len(coins); j++ {
			if i > coins[j] {
				count := dp[i-coins[j]] + 1
				if count < dp[i] {
					dp[i] = count
				}
			}
		}
	}
	if dp[amount] != 999999 {
		return dp[amount]
	}
	return -1
}
