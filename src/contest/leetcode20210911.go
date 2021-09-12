package contest

import "strconv"

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

func InterchangeableRectangles(rectangles [][]int) int64 {
	var res int64 = 0
	numMap := make(map[string]int)
	for i := 0; i < len(rectangles); i++ {
		fcb := divisor(rectangles[i][0], rectangles[i][1])
		tmp1 := rectangles[i][0] / fcb
		tmp2 := rectangles[i][1] / fcb
		numMap[strconv.Itoa(tmp1)+"/"+strconv.Itoa(tmp2)]++
	}
	for _, v := range numMap {
		res += int64(v * (v - 1) / 2)
	}
	return res
}

func divisor(i, j int) (maxDivisor int) {
	max, min := i, j
	if j > i {
		max = j
		min = i
	}
	//用小数除大数取余
	complement := max % min
	//余数不为零，将余数作为小数，除数作为大数，递归调用自己
	if complement != 0 {
		maxDivisor = divisor(complement, min)
	} else {
		//余数为零，除数则为最大公约数
		maxDivisor = min
	}
	return
}
