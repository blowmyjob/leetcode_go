package leetcode

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	halfNum := 0
	for halfNum < x {
		halfNum = halfNum*10 + x%10
		x /= 10
	}
	return halfNum == x || halfNum == x/10
}
