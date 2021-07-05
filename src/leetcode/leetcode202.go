package leetcode

func IsHappy(n int) bool {
	happyMap := make(map[int]int, 10)
	for n != 1 {
		if _, ok := happyMap[n]; ok {
			return false
		}
		happyMap[n] = 1
		n = transferNum(n)
	}
	return true
}

func transferNum(num int) int {
	sum := 0
	for num != 0 {
		var temp int = num % 10
		sum += temp * temp
		num /= 10
	}
	return sum
}
