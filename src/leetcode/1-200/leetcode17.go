package leetcode

var word = map[int]string{
	0: "",
	1: "",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func LetterCombinations(digits string) []string {
	res := []string{}
	if len(digits) == 0 {
		return res
	}
	var dfsLetter func(index int, str, digits string)
	dfsLetter = func(index int, str, digits string) {
		if index == len(digits) {
			res = append(res, str)
			return
		}
		tmpIndex := digits[index] - '0'
		tmpStr := word[int(tmpIndex)]
		for i := 0; i < len(tmpStr); i++ {
			str = str + string(tmpStr[i])
			dfsLetter(index+1, str, digits)
			str = str[0 : len(str)-1]
		}
	}
	dfsLetter(0, "", digits)
	return res
}
