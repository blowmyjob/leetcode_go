package leetcode

func maxProfit(prices []int) int {
	buy1, buy2 := -prices[0], -prices[0]
	sell1, sell2 := 0, 0
	for i := 0; i < len(prices); i++ {
		buy1 = MaxPrices(buy1, -prices[i])
		sell1 = MaxPrices(sell1, buy1+prices[i])
		buy2 = MaxPrices(buy2, sell1-prices[i])
		sell2 = MaxPrices(sell2, buy2+prices[i])
	}
	return sell2
}

func MaxPrices(i, j int) int {
	if i > j {
		return i
	}
	return j
}
