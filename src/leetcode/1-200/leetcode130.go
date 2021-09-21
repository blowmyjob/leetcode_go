package leetcode

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1 {
				dfsSolve(board, i, j)
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '1' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func dfsSolve(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = '1'
	dfsSolve(board, i-1, j)
	dfsSolve(board, i+1, j)
	dfsSolve(board, i, j-1)
	dfsSolve(board, i, j+1)
}
