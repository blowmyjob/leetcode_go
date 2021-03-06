package leetcode

func grayCode(n int) []int {
	res := []int{}
	for i := 0; i < 1<<n; i++ {
		res = append(res, i^i>>1)
	}
	return res
}
