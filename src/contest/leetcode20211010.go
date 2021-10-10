package contest

import (
	"math"
	"sort"
)

func minOperations1(grid [][]int, x int) int {
	res := 0
	arr := []int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j]%x != grid[0][0]%x {
				return -1
			}
			arr = append(arr, grid[i][j]/x)
		}
	}
	sort.Ints(arr)
	mid := len(arr) / 2
	for i := 0; i < len(arr); i++ {
		res += int(math.Abs(float64(arr[mid] - arr[i])))
	}
	return res
}
