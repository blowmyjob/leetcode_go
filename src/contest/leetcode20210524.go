package contest

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	var res int = 0
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		if i == 0 && arr[0] != 1 {
			arr[i] = 1
			res = max(res, arr[i])
		} else if i > 0 {
			arr[i] = min(arr[i-1], arr[i-1]+1)
			res = max(res, arr[i])
		}
	}
	return res
}

func CheckZeroOnes(s string) bool {
	var res0 int = 0
	var res1 int = 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			tmpLen := 1
			for j := i + 1; j < len(s); j++ {
				if s[j-1] == s[j] && s[j] == '0' {
					tmpLen++
				} else {
					break
				}
			}
			if tmpLen > res0 {
				res0 = tmpLen
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			tmpLen := 1
			for j := i + 1; j < len(s); j++ {
				if s[j-1] == s[j] && s[j] == '1' {
					tmpLen++
				} else {
					break
				}
			}
			if tmpLen > res1 {
				res1 = tmpLen
			}
		}
	}
	return res0 < res1
}
