package leetcode

func toHex(num int) string {
	s := "0123456789abcdef"
	if num == 0 {
		return "0"
	}
	n := int64(num)
	if num < 0 {
		n = int64(num) + 1<<32
	}
	str := ""
	for n != 0 {
		str = string(s[n%16]) + str
		n /= 16
	}
	return str
}
