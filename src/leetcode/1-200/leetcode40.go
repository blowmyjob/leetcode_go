package leetcode

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	arr := []int{}
	dfsCombs(candidates, arr, &res, 0, target)
	return res
}

func dfsCombs(candidates []int, arr []int, res *[][]int, index int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		*res = append(*res, tmp)
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		if i > index && candidates[i-1] == candidates[i] {
			continue
		}
		arr = append(arr, candidates[i])
		dfsCombs(candidates, arr, res, i+1, target-candidates[i])
		arr = arr[:len(arr)-1]
	}
}
