package leetcode

func IsAdditiveNumber(num string) bool {
	var helper func(num string, index, len int, pre, sum int64, k int) bool
	helper = func(num string, index, len int, pre, sum int64, k int) bool {
		if len == index {
			return k > 2
		}
		for i := index; i < len; i++ {
			curNum := sumNum(num, index, i)
			if curNum < 0 {
				continue
			}
			if k >= 2 && curNum != sum {
				continue
			}
			if helper(num, i+1, len, curNum, curNum+pre, k+1) {
				return true
			}
		}
		return false
	}
	return helper(num, 0, len(num), 0, 0, 0)
}

func sumNum(num string, l, r int) int64 {
	if l < r && num[l] == '0' {
		return -1
	}
	var res int64 = 0
	for i := l; i <= r; i++ {
		res = res*10 + int64(num[i]-'0')
	}
	return res
}
