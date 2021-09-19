package leetcode

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber != 0 {
		columnNumber--
		tmp := byte('A' + (columnNumber % 26))
		columnNumber /= 26
		res = string(tmp) + res
	}
	return res
}
