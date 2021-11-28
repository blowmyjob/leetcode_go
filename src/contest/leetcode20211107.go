package contest

import (
	"sort"
	"strings"
)

func countVowels(word string) (ans int64) {
	for i, ch := range word {
		if strings.ContainsRune("aeiou", ch) {
			ans += int64(i+1) * int64(len(word)-i)
		}
	}
	return
}

func countVowelSubstrings(word string) int {
	var sum int = 0
	var wordCount func(word string) int
	wordCount = func(word string) int {
		num1 := strings.Count(word, "a")
		num2 := strings.Count(word, "e")
		num3 := strings.Count(word, "i")
		num4 := strings.Count(word, "o")
		num5 := strings.Count(word, "u")
		if num1 <= 0 || num2 <= 0 || num3 <= 0 || num4 <= 0 || num5 <= 0 {
			return 0
		}
		if num1+num2+num3+num4+num5 != len(word) {
			return 0
		}
		return num1 + num2 + num3 + num4 + num5
	}
	for i := 0; i < len(word); i++ {
		for j := i + 1; j < len(word)+1; j++ {
			subStr := word[i:j]
			sum += wordCount(subStr)
		}
	}
	return sum
}

func minimizedMaximum(n int, quantities []int) int {
	return sort.Search(1e5, func(max int) bool {
		cnt := 0
		for _, q := range quantities {
			cnt += (q + max) / (max + 1)
		}
		return cnt <= n
	}) + 1
}
