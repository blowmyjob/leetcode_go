package contest

func largestAltitude(gain []int) int {
	temp, max := 0, 0
	for i := 0; i < len(gain); i++ {
		temp += gain[i]
		if temp > max {
			max = temp
		}
	}
	return max
}

func Decode(encoded []int) []int {
	a := 0
	for i := len(encoded) + 1; i >= 1; i-- {
		a = a ^ i
	}
	for i := len(encoded) - 1; i >= 1; i -= 2 {
		a = a ^ encoded[i]
	}
	ans := make([]int, len(encoded)+1)
	ans[0] = a
	for i := 1; i <= len(encoded); i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}
	return ans
}
