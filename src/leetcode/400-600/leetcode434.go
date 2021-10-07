package leetcode

func countSegments(s string) int {
	ans := 0
	for i, ch := range s {
		if (i == 0 || s[i-1] == ' ') && ch != ' ' {
			ans++
		}
	}
	return ans
}
