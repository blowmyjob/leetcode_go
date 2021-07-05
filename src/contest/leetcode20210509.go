package contest

import (
	"sort"
	"strconv"
)

func minSwaps(s string) int {
	var i0 int = 1
	var i1 int = 1
	var res int = 0
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			if s[i] == 1 {
				for i0 < len(s) {
					if s[i0] == s[i0-1] && s[i0] == 0 {
						res++
						break
					}
					i0++
				}
			} else {
				for i1 < len(s) {
					if s[i1] == s[i1-1] && s[i1] == 1 {
						res++
						break
					}
					i1++
				}
			}
		}
	}
	return res
}

func minPairSum(nums []int) int {
	var res = 0
	var left = 0
	var right = len(nums) - 1
	sort.Ints(nums)
	for left < right {
		if nums[left]+nums[right] > res {
			res = nums[left] + nums[right]
		}
		left++
		right--
	}
	return res
}

func MinimumXORSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	numMap1 := make(map[int]int)
	numMap2 := make(map[int]int)
	set1 := make([]int, 0)
	set2 := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		numMap1[nums1[i]]++
	}
	for i := 0; i < len(nums2); i++ {
		numMap2[nums2[i]]++
	}
	for _, v := range nums2 {
		time := numMap1[v]
		if time == 0 {
			set1 = append(set1, v)
		} else {
			numMap1[v]--
		}
	}
	for _, v := range nums1 {
		time := numMap2[v]
		if time == 0 {
			set2 = append(set2, v)
		} else {
			numMap2[v]--
		}
	}
	sort.Ints(set1)
	sort.Ints(set2)
	var res int = 0
	for i := 0; i < len(set1); i++ {
		tmp := set1[i] ^ set2[i]
		res += tmp
	}
	return res
}

func MaxValue(n string, x int) string {
	if n[0] == '-' {
		for i := 1; i < len(n); i++ {
			tmp, _ := strconv.Atoi(string(n[i]))
			if x < tmp {
				n := n[:i] + strconv.Itoa(x) + n[i:]
				return n
			}
		}
		n := n[:len(n)] + strconv.Itoa(x) + n[len(n):]
		return n
	} else {
		for i := 0; i < len(n); i++ {
			tmp, _ := strconv.Atoi(string(n[i]))
			if x > tmp {
				if i > 0 {
					n := n[:i] + strconv.Itoa(x) + n[i:]
					return n
				} else {
					n := n[:0] + strconv.Itoa(x) + n[0:]
					return n
				}
			}
		}
		n := n[:len(n)] + strconv.Itoa(x) + n[len(n):]
		return n
	}
}

func assignTasks(servers []int, tasks []int) []int {
	res := make([]int, len(tasks))
	for i := 0; i < len(tasks); i++ {
		for j := 0; j < len(servers); j++ {

		}
	}
	return res
}

func ReductionOperations(nums []int) int {
	var res int = 0
	sort.Ints(nums)
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]]++
	}
	nums = Rm_duplicate(nums)
	for i := 1; i < len(nums); i++ {
		min := nums[0]
		max := nums[i]
		if min == max {
			continue
		}
		res += numMap[nums[i]] * (i)
	}
	return res
}

func Rm_duplicate(list []int) []int {
	var x []int = []int{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}
