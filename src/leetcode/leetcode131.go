package leetcode

func Partition(s string) [][]string {
	res := [][]string{}
	arr := []string{}
	var helper func(s string, arr []string, index int)
	helper = func(s string, arr []string, index int) {
		if index == len(s) {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		for i := index + 1; i < len(s); i++ {
			str := s[index:i]
			if isPalindrome2(str) {
				arr = append(arr, str)
				helper(s, arr, i+1)
				arr = arr[:len(arr)-1]
			}
		}
	}
	helper(s, arr, 0)
	return res
}

func isPalindrome2(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
