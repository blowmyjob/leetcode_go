package leetcode

import "strconv"

func LargestOddNumber(num string) string {
	var max int64 = 0
	for i := 0; i < len(num); i++ {
		for j := i; j <= len(num); j++ {
			temp, _ := strconv.ParseInt(num[i:j], 10, 64)
			if temp%2 == 1 && temp > max {
				max = temp
			}
		}
	}
	if max != 0 {
		str := strconv.FormatInt(max, 10)
		return str
	}
	return ""
}
