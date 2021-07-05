package leetcode

func restoreArray(adjacentPairs [][]int) []int {
	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}
	res := make([]int, 0)
	return res
}

func countBalls(lowLimit int, highLimit int) int {
	max := 0
	numMap := make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		order := calNums(i)
		numMap[order]++
	}
	for i := lowLimit; i <= highLimit; i++ {
		order := calNums(i)
		max = maxNum(numMap[order], max)
	}
	return max
}
func maxNum(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
func calNums(n int) int {
	res := 0
	for n != 0 {
		res += n % 10
		n /= 10
	}
	return res
}
