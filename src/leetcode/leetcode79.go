package leetcode

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if helper(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func helper(board [][]byte, row, col int, word string, index int) bool {
	if index >= len(word) {
		return true
	}
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) || word[index] != board[row][col] {
		return false
	}
	board[row][col] ^= 255
	flag := helper(board, row+1, col, word, index+1) ||
		helper(board, row-1, col, word, index+1) ||
		helper(board, row, col+1, word, index+1) ||
		helper(board, row, col-1, word, index+1)
	board[row][col] ^= 255
	return flag
}
