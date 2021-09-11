package contest

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	res := 0
	wordMap := make(map[int]int)
	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[i]); j++ {
			wordMap[source[i][j]] = 1
		}
	}
	for i := 0; i < len(target); i++ {
		for j := 0; j < len(target[i]); j++ {
			if _, ok := wordMap[target[i][j]]; !ok {
				res++
			}
		}
	}
	return res
}

func MaxmiumScore(cards []int, cnt int) int {
	res := 0
	var helper func(card []int, index int, num int, sum int)
	helper = func(card []int, index int, num int, sum int) {
		if num == cnt {
			if sum%2 == 0 {
				if sum > res {
					res = sum
				}
			}
		}
		for i := index; i < len(cards); i++ {
			helper(cards, i+1, num+1, sum+cards[i])
		}
	}
	helper(cards, 0, 0, 0)
	return res
}
