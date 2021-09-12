package leetcode

func subsets(nums []int) [][]int {
	res := [][]int{}
	arr := []int{}
	var helper func(nums, arr []int, index int)
	helper = func(nums, arr []int, index int) {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		res = append(res, tmp)
		for i := index; i < len(nums); i++ {
			arr = append(arr, nums[i])
			helper(nums, arr, i+1)
			arr = arr[:len(arr)-1]
		}
	}
	helper(nums, arr, 0)
	return res
}
