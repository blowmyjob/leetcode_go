package leetcode

func romanToInt(s string) int {
	count := 0
	wordMap := make(map[string]int)
	wordMap["M"] = 1000
	wordMap["CM"] = 900
	wordMap["D"] = 500
	wordMap["CD"] = 400
	wordMap["C"] = 100
	wordMap["XC"] = 90
	wordMap["L"] = 50
	wordMap["XL"] = 40
	wordMap["X"] = 10
	wordMap["IX"] = 9
	wordMap["V"] = 5
	wordMap["IV"] = 4
	wordMap["I"] = 1
	for i := 0; i < len(s); {
		if i+1 < len(s) && wordMap[s[i:i+2]] != 0 {
			count += wordMap[s[i:i+2]]
			i += 2
		} else if wordMap[s[i:i+1]] != 0 {
			count += wordMap[s[i:i+1]]
			i += 1
		}
	}
	return count
}
