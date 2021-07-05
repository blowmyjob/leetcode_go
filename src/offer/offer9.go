package offer

import (
	"sort"
	"strconv"
)

func GetLeastNumbers_Solution(input []int, k int) []int {
	if k > len(input) {
		return nil
	}
	sort.Ints(input)
	return input[:k]
}

func FindGreatestSumOfSubArray(array []int) int {
	n := len(array)
	if n == 0 {
		return 0
	}
	max, sum := array[0], 0
	for i := 0; i < n; i++ {
		if sum > 0 {
			sum += array[i]
		} else {
			sum = array[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func PrintMinNumber(numbers []int) string {
	// write code here
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum1, _ := strconv.Atoi(strconv.Itoa(numbers[i]) + strconv.Itoa(numbers[j]))
			sum2, _ := strconv.Atoi(strconv.Itoa(numbers[j]) + strconv.Itoa(numbers[i]))
			if sum1 > sum2 {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	res := ""
	for i := 0; i < n; i++ {
		res += strconv.Itoa(numbers[i])
	}
	return res
}

func FindNumbersWithSum(array []int, sum int) []int {
	// write code here
	i, j := 0, len(array)-1
	for i < j {
		if array[i]+array[j] == sum {
			return []int{array[i], array[j]}
		} else if array[i]+array[j] > sum {
			i++
		} else {
			j++
		}
	}
	return nil
}

func Duplicate(numbers []int, duplication *[1]int) bool {
	numMap := make(map[int]int, 0)
	for i := 0; i < len(numbers); i++ {
		if _, ok := numMap[numbers[i]]; ok {
			//存在
			duplication[0] = numbers[i]
			return true
		}
		numMap[numbers[i]] = 1
	}
	return false
}
