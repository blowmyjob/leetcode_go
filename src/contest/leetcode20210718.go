package contest

import "math"

func addRungs(rungs []int, dist int) int {
	count := 0
	startIndex := 0
	currentHeight := 0
	for startIndex < len(rungs) {
		if currentHeight+dist < rungs[startIndex] {
			add := (rungs[startIndex] - currentHeight) / dist
			currentHeight += add * dist
			if currentHeight >= rungs[startIndex] {
				add--
			}
			count += add
		} else {
			currentHeight = rungs[startIndex]
			startIndex++
		}
	}
	return count
}

func maxPoints(points [][]int) int64 {
	dp := [][]int{}
	row, col := len(points), len(points[0])
	for i := 0; i < row; i++ {
		inDp := make([]int, col)
		dp = append(dp, inDp)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 {
				dp[i][j] = points[i][j]
			} else {
				dp[i][j] = math.MinInt8
				for k := 0; k < col; k++ {
					abs := ((k - j) ^ (k-j)>>31) - (k-j)>>31
					tmp := dp[i-1][k] + points[i][j] - abs
					if tmp > dp[i][j] {
						dp[i][j] = tmp
					}
				}
			}
		}
	}
	var res int64 = math.MinInt64
	for i := 0; i < col; i++ {
		if int64(dp[row-1][i]) > res {
			res = int64(dp[row-1][i])
		}
	}
	return res
}
