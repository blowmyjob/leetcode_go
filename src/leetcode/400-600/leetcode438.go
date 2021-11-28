package leetcode

import (
	"strings"
)

func FindAnagrams(s string, p string) []int {
	res := []int{}
	funcAnagrams := func(s, p string) bool {
		for i := 0; i < len(s); i++ {
			str := string(s[i])
			count1 := strings.Count(s, str)
			count2 := strings.Count(p, str)
			if count1 != count2 {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(s)-len(p); i++ {
		str := s[i : i+len(p)]
		if funcAnagrams(str, p) {
			res = append(res, i)
		}
	}
	return res
}

func GetAverages(nums []int, k int) []int {
	sum := make([]int, len(nums)+1)
	res := make([]int, len(nums))
	if k == 0 {
		return nums
	}
	if k > len(nums) {
		return []int{-1}
	}
	for i := 1; i <= len(nums); i++ {
		sum[i] += sum[i-1] + nums[i-1]
	}
	for i := 0; i < k; i++ {
		res[i] = -1
	}
	for i := len(nums) - 1; i >= len(nums)-k; i-- {
		res[i] = -1
	}
	for i := k; i < len(nums)-k; i++ {
		right := i + k
		left := i - k
		sums := sum[right] - sum[left]
		res[i] = sums / (2*k + 1)
	}
	return res
}

func minimumDeletions(nums []int) int {
	i, j := 0, 0
	for p, num := range nums {
		if num < nums[i] {
			i = p
		}
		if num > nums[j] {
			j = p
		}
	}
	if i > j {
		i, j = j, i // 保证 i <= j
	}
	return min(min(j+1, len(nums)-i), i+1+len(nums)-j)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
