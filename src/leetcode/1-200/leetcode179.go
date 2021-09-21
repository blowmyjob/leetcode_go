package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

func LargestNumber(nums []int) string {
	n := len(nums)
	strs := make([]string, n)
	for i, _ := range nums {
		strs[i] = fmt.Sprintf("%d", nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	if strs[0][0] == '0' {
		return "0"
	}
	return strings.Join(strs, "")
}
