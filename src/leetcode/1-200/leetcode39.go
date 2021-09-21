package leetcode

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	tmp := []int{}
	sort.Ints(candidates)
	dfsComb(candidates, target, 0, tmp, &res)
	return res
}

func dfsComb(candidates []int, target int, index int, tmp []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		tmpArr := make([]int, len(tmp))
		copy(tmpArr, tmp)
		*res = append(*res, tmpArr)
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target { // 这里可以剪枝优化
			break
		}
		tmp = append(tmp, candidates[i])
		dfsComb(candidates, target-candidates[i], i, tmp, res)
		tmp = tmp[:len(tmp)-1]
	}
}
