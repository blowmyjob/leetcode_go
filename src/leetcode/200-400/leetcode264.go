package leetcode

func NthUglyNumber(n int) int {
	u2, u3, u5 := 1, 1, 1
	res := make([]int, n+1)
	res[0], res[1] = 0, 1
	for i := 2; i <= n; i++ {
		res[i] = minUgly(res[u2]*2, minUgly(res[u3]*3, res[u5]*5))
		if res[i]%2 == 0 {
			u2++
		}
		if res[i]%3 == 0 {
			u3++
		}
		if res[i]%5 == 0 {
			u5++
		}
	}
	return res[n]
}

func minUgly(i, j int) int {
	if i > j {
		return j
	}
	return i
}
