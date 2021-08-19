package contest

import (
	"sort"
	"strings"
)

func RearrangeArray(nums []int) []int {
	sort.Ints(nums)
	res := []int{}
	left, right := 0, len(nums)/2+1
	for left <= len(nums)/2 || right < len(nums) {
		if left <= len(nums)/2 {
			res = append(res, nums[left])
		}
		if right < len(nums) {
			res = append(res, nums[right])
		}
		left++
		right++
	}
	return res
}

func numOfStrings(patterns []string, word string) int {
	res := 0
	for _, v := range patterns {
		if strings.Contains(word, v) {
			res++
		}
	}
	return res
}
