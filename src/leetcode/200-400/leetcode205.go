package leetcode

func isIsomorphic(s string, t string) bool {
	wordMap1 := make(map[byte]byte)
	wordMap2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		s1, ok1 := wordMap1[s[i]]
		s2, ok2 := wordMap2[t[i]]
		if ok1 && ok2 {
			if s1 != t[i] && s2 != s[i] {
				return false
			}
		} else if !ok1 && !ok2 {
			wordMap1[s[i]] = t[i]
			wordMap2[t[i]] = s[i]
		} else {
			return false
		}
	}
	return true
}
