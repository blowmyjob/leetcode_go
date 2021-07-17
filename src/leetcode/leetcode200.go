package leetcode

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfsIsland(grid, i, j)
				count++
			}
		}
	}
	return count
}

func dfsIsland(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	dfsIsland(grid, i-1, j)
	dfsIsland(grid, i+1, j)
	dfsIsland(grid, i, j-1)
	dfsIsland(grid, i, j+1)
}
