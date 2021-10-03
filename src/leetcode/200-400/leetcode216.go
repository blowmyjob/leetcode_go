package leetcode

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	arr := []int{}
	var helper func(k int, n int, index int, arr []int)
	helper = func(k int, n int, index int, arr []int) {
		if len(arr) == k && n == 0 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		}
		for i := index; i <= 9; i++ {
			arr = append(arr, i)
			helper(k, n-i, i+1, arr)
			arr = arr[:len(arr)-1]
		}
	}
	helper(k, n, 1, arr)
	return res
}
