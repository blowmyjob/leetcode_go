package leetcode

func integerBreak(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	res := make([]int, n+1)
	res[1] = 1
	res[2] = 2
	res[3] = 3
	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			if res[i-j]*res[j] > res[i] {
				res[i] = res[i-j] * res[j]
			}
		}
	}
	return res[n]
}
