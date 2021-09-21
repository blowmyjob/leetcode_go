package leetcode

func findRepeatedDnaSequences(s string) []string {
	wordMap := make(map[string]int)
	setMap := make(map[string]int)
	strs := []string{}
	for i := 0; i <= len(s)-10; i++ {
		tmp := s[i : i+10]
		wordMap[tmp]++
		if wordMap[tmp] > 1 {
			strs = append(strs, tmp)
			setMap[tmp] = 1
		}
	}
	return strs
}
