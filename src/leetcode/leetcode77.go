package leetcode

func Combine(n int, k int) [][]int {
	res := [][]int{}
	arr := []int{}
	var helper func(arr []int, index, n int)
	helper = func(arr []int, index, n int) {
		if len(arr) == k {
			tmp := make([]int, k)
			copy(tmp, arr)
			res = append(res, tmp)
		}
		for i := index; i <= n; i++ {
			arr = append(arr, i)
			helper(arr, i+1, n)
			arr = arr[:len(arr)-1]
		}
	}
	helper(arr, 1, n)
	return res
}
