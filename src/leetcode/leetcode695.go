package leetcode

func maxAreaOfIsland(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				tmp := dfsMaxIsland(grid, i, j)
				if tmp > count {
					count = tmp
				}
			}
		}
	}
	return count
}

func dfsMaxIsland(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0
	return 1 + dfsMaxIsland(grid, i-1, j) + dfsMaxIsland(grid, i+1, j) + dfsMaxIsland(grid, i, j-1) + dfsMaxIsland(grid, i, j+1)
}
