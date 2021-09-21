package contest

import (
	"math"
	"strings"
)

func removeOccurrences(s string, part string) string {
	for strings.Contains(s, part) {
		index := strings.Index(s, part)
		str := s[0:index] + s[index+len(part):]
		s = str
	}
	return s
}
func maxAlternatingSum(nums []int) int64 {
	f := [2]int{0, math.MinInt64 / 2}
	for _, v := range nums {
		f = [2]int{max(f[0], f[1]-v), max(f[0]+v, f[1])}
	}
	return int64(f[1])
}

func WonderfulSubstrings(word string) int64 {
	var res int64 = 0
	numMap := make(map[string]int)
	for i := 0; i < len(word); i++ {
		wordMap := make(map[string]int)
		for j := i; j < len(word); j++ {
			wordMap[string(word[i])]++
			s := word[i : j+1]
			_, ok := numMap[s]
			if isWanMei(wordMap) && !ok {
				res++
				numMap[s]++
			}
		}
	}
	return res
}
func isWanMei(wordMap map[string]int) bool {
	var nums int = 0
	for _, v := range wordMap {
		if v%2 != 1 {
			nums++
		}
	}
	return nums <= 1
}
