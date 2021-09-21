package leetcode

import "sort"

func findLongestWord(s string, dictionary []string) string {
	var temp string
	sort.Strings(dictionary)
	for i := 0; i < len(dictionary); i++ {
		if find(s, dictionary[i]) && len(temp) < len(dictionary[i]) {
			temp = dictionary[i]
		}
	}
	return temp
}

func find(x, y string) bool {
	j := 0
	for i := 0; j < len(y) && i < len(x); i++ {
		if y[j] == x[i] {
			j++
		}
	}
	return j == len(y)
}
