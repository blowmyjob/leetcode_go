package leetcode

func maxProfit(prices []int) int {
	buy, sell := -prices[0], 0
	pre := -prices[0]
	var maxFunc func(i, j int) int
	maxFunc = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	for i := 0; i < len(prices); i++ {
		tmp := sell
		buy = maxFunc(buy, pre-prices[i])
		sell = maxFunc(sell, buy+prices[i])
		pre = tmp
	}
	return sell
}
