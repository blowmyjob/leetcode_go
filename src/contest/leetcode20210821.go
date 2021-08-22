package contest

import (
	"math"
	"sort"
	"strconv"
)

func maxMatrixSum(matrix [][]int) int64 {
	row, col := len(matrix), len(matrix[0])
	mark := false
	min := math.MaxInt64
	var res int64 = 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			x := matrix[i][j]
			if x < 0 {
				mark = !mark
				x = -1 * x
			}
			if min > x {
				min = x
			}
			res += int64(x)
		}
	}
	if mark {
		res -= 2 * int64(min)
	}
	return res
}

func FindDifferentBinaryString(nums []string) string {
	sort.Strings(nums)
	wordMap := make(map[string]int)
	for _, v := range nums {
		wordMap[v] = 1
	}
	if len(nums) == 1 && nums[0] == "1" {
		return "0"
	}
	if len(nums) == 1 && nums[0] == "0" {
		return "1"
	}
	n := len(nums)
	for i := 1; i <= int(math.Pow(2, float64(n))); i++ {
		str := DecToBin(i, n)
		_, ok := wordMap[str]
		if !ok {
			return str
		}
	}
	return ""
}

// 原理：除2取模是最低位
func DecToBin(n, num int) string {
	result := ""

	if n == 0 {
		return "0"
	}

	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	len1 := len(result)
	for i := len1; i < num; i++ {
		result = "0" + result
	}
	return result
}
