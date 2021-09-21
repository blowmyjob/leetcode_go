package contest

func chalkReplacer(chalk []int, k int) int {
	var i int = 0
	var sum int = 0
	length := len(chalk)
	for i := 0; i < len(chalk); i++ {
		sum += chalk[i]
	}
	k %= sum
	for k >= 0 {
		k -= chalk[i%length]
		if k < 0 {
			break
		}
		i++
	}
	return i % length
}

func MaximumRemovals(s string, p string, removable []int) int {
	wordMap := make(map[int]int)
	for i := 0; i < len(s); i++ {
		wordMap[i] = 1
	}
	for i := 0; i < len(removable); i++ {
		wordMap[removable[i]] = 0
		str := ""
		for j := 0; j < len(s); j++ {
			if wordMap[j] == 0 {
				continue
			}
			str += string(s[j])
		}
		if !isSubsequence(p, str) {
			return i
		}
	}
	return 0
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	j := 0
	for i := 0; i < len(t); i++ {
		if t[i] == s[j] {
			j++
			if j == len(s) {
				return true
			}
		}
	}
	return false
}
