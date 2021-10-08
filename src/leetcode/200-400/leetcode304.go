package leetcode

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var numMatrix NumMatrix
	m := len(matrix)
	if m == 0 {
		return numMatrix
	}
	n := len(matrix[0])
	numMatrix.sums = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		numMatrix.sums[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			numMatrix.sums[i][j] = numMatrix.sums[i-1][j] +
				numMatrix.sums[i][j-1] -
				numMatrix.sums[i-1][j-1] +
				matrix[i-1][j-1]
		}
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] -
		this.sums[row1][col2+1] -
		this.sums[row2+1][col1] + this.sums[row1][col1]
}
