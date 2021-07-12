package leetcode

func CountPalindromicSubsequence(s string) int {
	slice := [26][2]int{}
	ret := 0
	for i := 0; i < 26; i++ {
		slice[i][0] = -1
	}
	for i, v := range s {
		if slice[v-'a'][0] == -1 {
			slice[v-'a'][0] = i
		} else {
			slice[v-'a'][1] = i
		}
	}
	for _, v := range slice {
		set := [26]bool{}
		if v[0] != -1 {
			for i := v[0] + 1; i < v[1]; i++ {
				if !set[s[i]-'a'] {
					ret++
					set[s[i]-'a'] = true
				}
			}
		}
	}
	return ret
}
