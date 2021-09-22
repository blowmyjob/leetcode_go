package leetcode

import "sort"

func Permute(nums []int) [][]int {
	res := [][]int{}
	arr := []int{}
	sort.Ints(nums)
	var helper func(arr []int, index int)
	helper = func(arr []int, index int) {
		if index == len(nums) {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
			arr = append(arr, nums[i])
			nums[i], nums[index] = nums[index], nums[i]
			helper(arr, index+1)
			nums[i], nums[index] = nums[index], nums[i]
			arr = arr[:len(arr)-1]
		}
	}
	helper(arr, 0)
	return res
}
