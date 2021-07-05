package leetcode

func maxProfit2(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		temp := prices[i] - prices[i-1]
		if temp > 0 {
			profit += temp
		}
	}
	return profit
}
