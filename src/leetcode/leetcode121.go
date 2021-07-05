package leetcode

func maxProfit1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit, minPrice := 0, prices[0]
	for i := 0; i < len(prices); i++ {
		profit = maxProfitFuc(profit, prices[i]-minPrice)
		minPrice = min(prices[i], minPrice)
	}
	return profit
}

func min(x1, x2 int) int {
	if x1 > x2 {
		return x2
	}
	return x1
}

func maxProfitFuc(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}
