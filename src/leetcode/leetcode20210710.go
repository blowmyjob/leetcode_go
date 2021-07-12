package leetcode

import "fmt"

func NearestExit(maze [][]byte, entrance []int) int {
	queue := [][]int{}
	var visit [][]int
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue = append(queue, []int{entrance[0], entrance[1]})
	cnt := 1
	n, m := len(maze), len(maze[0])
	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		visit = append(visit, tmp)
	}
	visit[entrance[0]][entrance[1]] = 1
	for len(queue) > 0 {
		size := len(queue)
		for j := 0; j < size; j++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			for i := 0; i < len(direct); i++ {
				rowi := row + direct[i][0]
				coli := col + direct[i][1]
				if rowi >= 0 && rowi < n && coli >= 0 && coli < m && maze[rowi][coli] == '.' && visit[rowi][coli] != 1 {
					fmt.Printf("%d_%d\n", rowi, coli)
					if rowi == 0 || rowi == n-1 || coli == m-1 || coli == 0 {
						return cnt
					}
					visit[rowi][coli] = 1
					queue = append(queue, []int{rowi, coli})
				}
			}
		}
		cnt++
	}
	return -1
}

func countTriples(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if i*i+j*j == k*k {
					ans++
				}
			}
		}
	}
	return ans
}
