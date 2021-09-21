package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	leftRow, rightCol := 0, len(matrix[0])-1
	for leftRow < len(matrix) && rightCol >= 0 {
		if matrix[leftRow][rightCol] > target {
			rightCol--
		} else if matrix[leftRow][rightCol] < target {
			leftRow++
		} else {
			return true
		}
	}
	return false
}
