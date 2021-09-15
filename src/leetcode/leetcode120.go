package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) <= 0 {
		return 0
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = minSum(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}

func minSum(i, j int) int {
	if i < j {
		return i
	}
	return j
}
