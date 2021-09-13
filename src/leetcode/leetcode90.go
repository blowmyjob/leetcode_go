package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	arr := []int{}
	sort.Ints(nums)
	var helper func(arr []int, index int, res *[][]int)
	helper = func(arr []int, index int, res *[][]int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			arr = append(arr, nums[i])
			helper(arr, i+1, res)
			arr = arr[:len(arr)-1]
		}
	}
	helper(arr, 0, &res)
	return res
}
