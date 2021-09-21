package leetcode

import "strings"

func IsPalindrome(s string) bool {
	var str string
	for _, v := range s {
		if isRune(v) {
			str += string(v)
		}
	}
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func isRune(char rune) bool {
	return (char >= '0') && (char <= '9') || ((char >= 'a') && (char <= 'z')) || ((char >= 'A') && (char <= 'Z'))
}
