package leetcode

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	str := make([]int, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			str[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	for i := m + n - 1; i > 0; i-- {
		str[i-1] += str[i] / 10
		str[i] %= 10
	}
	if str[0] == 0 {
		str = str[1:]
	}
	res := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		res[i] = '0' + byte(str[i])
	}
	return string(res)
}
