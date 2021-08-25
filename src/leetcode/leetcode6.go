package leetcode

func reverse(x int) int {
	num := x
	if x < 0 {
		num *= -1
	}
	count := 0
	for num != 0 {
		count = count*10 + num%10
		num /= 10
	}
	if x < 0 {
		count *= -1
	}
	if count > 1<<31-1 || count < -(1<<31) {
		return 0
	}
	return count
}
