package leetcode

import "math"

func MyAtoi(s string) int {
	n := len(s)
	index := 0
	for index < n {
		if s[index] == ' ' {
			index++
		} else {
			break
		}
	}
	if index == n {
		return 0
	}
	negative := false
	if s[index] == '-' {
		negative = true
		index++
	} else if s[index] == '+' {
		negative = false
		index++
	} else if s[index] < '0' || s[index] > '9' {
		return 0
	}
	x := 0
	for index < n && (s[index] >= '0' && s[index] <= '9') {
		num := int(s[index] - '0')
		if x*10+num > 1<<31-1 {
			if negative {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		x = x*10 + num
		index++
	}
	if negative {
		x *= -1
	}
	return x
}
