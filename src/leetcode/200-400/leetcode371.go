package leetcode

func getSum(a int, b int) int {
	for b != 0 {
		temp := a ^ b
		carry := (a & b) << 1
		a = temp
		b = carry
	}
	return a
}
