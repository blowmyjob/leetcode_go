package leetcode

import "strings"

func reverseWords(s string) string {
	ss := strings.Fields(s)
	for i := 0; i <= len(ss)/2; i++ {
		ss[i], ss[len(ss)-i-1] = ss[len(ss)-i-1], ss[i]
	}
	return strings.Join(ss, "")
}
