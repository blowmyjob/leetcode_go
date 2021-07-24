package leetcode

func closedIsland(grid [][]int) int {
	num, rowNum, colNum := 0, len(grid), len(grid[0])
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 == 1 {
				continue
			}
			// 如果满足四面环水，则孤岛数量+1
			if dfsClosed(grid, rowNum, colNum, row, col) {
				num++
			}
		}
	}

	return num
}

func dfsClosed(grid [][]int, rowNum, colNum, row int, col int) (b bool) {
	// 如果触碰到边界了，无法构成孤岛
	if row < 0 || col < 0 || row >= rowNum || col >= colNum {
		return false
	}
	// 如果碰到水了，返回正确
	if grid[row][col] == 1 {
		return true
	}
	// 经过的地方置为海洋
	grid[row][col] = 1
	// 需要四个方向都环水
	up := dfsClosed(grid, rowNum, colNum, row-1, col)
	down := dfsClosed(grid, rowNum, colNum, row+1, col)
	left := dfsClosed(grid, rowNum, colNum, row, col-1)
	right := dfsClosed(grid, rowNum, colNum, row, col+1)
	return up && down && left && right
}
