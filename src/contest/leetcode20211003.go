package contest

func missingRolls(rolls []int, mean int, n int) []int {
	rollSum := 0
	for _, v := range rolls {
		rollSum += v
	}
	allSum := (len(rolls) + n) * mean
	nSum := allSum - rollSum
	res := [][]int{}
	var helper func(k int, n int, index int, arr []int)
	helper = func(k int, n int, index int, arr []int) {
		if n < 0 {
			return
		}
		if len(arr) == k && n == 0 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}
		for i := index; i <= 6; i++ {
			arr = append(arr, i)
			helper(k, n-i, i, arr)
			arr = arr[:len(arr)-1]
		}
	}
	helper(n, nSum, 1, []int{})
	if len(res) == 0 {
		return []int{}
	}
	return res[0]
}
