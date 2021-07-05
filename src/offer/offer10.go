package offer

func GetUglyNumber_Solution(index int) int {
	if index == 1 {
		return 1
	}
	dp := make([]int, index+1)
	a, b, c := 0, 0, 0
	for i := 2; i <= index; i++ {
		dp[i] = minNums(dp[a]*2, dp[b]*3, dp[c]*5)
		if dp[i]%2 == 0 {
			a++
		}
		if dp[i]%3 == 0 {
			b++
		}
		if dp[i]%5 == 0 {
			c++
		}
	}
	return dp[index]
}

func minNums(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else if c <= a && c <= b {
		return c
	}
	return 0
}
