package leetcode

func lengthOfLongestSubstring(s string) int {
	wordMap := make(map[byte]int)
	ans, right := 0, -1
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(wordMap, s[i-1])
		}
		for right+1 < len(s) && wordMap[s[right+1]] == 0 {
			wordMap[s[right+1]]++
			right++
		}
		ans = maxAns(ans, right-i+1)
	}
	return ans
}

func maxAns(x, y int) int {
	if x > y {
		return x
	}
	return y
}
