package leetcode

import "strings"

func WordPattern(pattern string, s string) bool {
	wordMap := make(map[byte]string)
	stop := make(map[string]byte)
	strs := strings.Split(s, " ")
	if len(strs) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		v, ok := wordMap[pattern[i]]
		_, sOk := stop[strs[i]]
		if ok && v != strs[i] {
			return false
		} else if !ok && !sOk {
			wordMap[pattern[i]] = strs[i]
			stop[strs[i]] = pattern[i]
		} else if (!ok && sOk) || (ok && !sOk) {
			return false
		}
	}
	return true
}
