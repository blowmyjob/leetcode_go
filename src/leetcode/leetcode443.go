package leetcode

import (
	"strconv"
)

func Compress(chars []byte) int {
	count, ctop, top := 1, 0, 0
	for i := 1; i < len(chars); i++ {
		if chars[i] != chars[ctop] {
			top++
			chars[top] = chars[i]
			ctop = top
			count = 1
		} else {
			count++
			if count > 1 {
				top = ctop
				str := strconv.Itoa(count)
				for j := 0; j < len(str); j++ {
					top++
					chars[top] = str[j]
				}
			}
		}
	}
	return top + 1
}
