package contest

func CheckMove(board [][]byte, rMove int, cMove int, color byte) bool {
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for i := 0; i < len(direct); i++ {
		dx := direct[i][0] + rMove
		dy := direct[i][1] + cMove
		if dx >= 0 && dx < 8 && dy >= 0 && dy < 8 && board[dx][dy] == color {
			continue
		}
		for dx >= 0 && dx < 8 && dy >= 0 && dy < 8 {
			if board[dx][dy] == '.' {
				break
			}
			if board[dx][dy] == color {
				return true
			}
			dx += direct[i][0]
			dy += direct[i][1]
		}
	}
	return false
}

func minSwaps1(s string) (ans int) {
	c := 0
	for _, b := range s {
		if b == '[' {
			c++
		} else if c > 0 {
			c--
		} else {
			c++ // 把最后面的 [ 和 ] 交换
			ans++
		}
	}
	return
}
