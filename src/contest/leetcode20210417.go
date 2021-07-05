package contest

import (
	"math"
	"sort"
)

func minOperations(nums []int) int {
	var res int = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		} else {
			res += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return res
}

func countPoints(points [][]int, queries [][]int) []int {
	resLen := len(queries)
	res := make([]int, resLen)
	for i := 0; i < len(queries); i++ {
		var num int = 0
		x := queries[i][0]
		y := queries[i][1]
		r := float64(queries[i][2])
		for j := 0; j < len(points); j++ {
			xLen := math.Abs(float64(points[j][0] - x))
			yLen := math.Abs(float64(points[j][1] - y))
			if xLen*xLen+yLen*yLen <= r*r {
				num++
			}
		}
		res[i] = num
	}
	return res
}

func GetMaximumXor(nums []int, maximumBit int) []int {
	var max int = 1
	res := make([]int, len(nums))
	for i := 1; i <= maximumBit; i++ {
		max *= 2
	}
	if max > 2 {
		max--
	}
	var xorValue int = nums[0]
	for i := 1; i < len(nums); i++ {
		xorValue ^= nums[i]
	}
	for i := len(nums) - 1; i >= 0; i-- {
		res[len(nums)-i-1] = xorValue ^ max
		xorValue ^= nums[i]
	}
	return res
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var res int = 0
	for i := 0; i < len(costs); i++ {
		if coins >= costs[i] {
			res++
			coins -= costs[i]
		} else {
			break
		}
	}
	return res
}

func getXORSum(arr1 []int, arr2 []int) int {
	arr := make([]int, 0)
	numMap := make(map[int]int, 0)
	var res int = 0
	if len(arr1) == 0 {
		arr = arr1
	} else if len(arr2) == 0 {
		arr = arr2
	}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			num := arr1[i] & arr2[j]
			numMap[num] = 1
			arr = append(arr, num)
		}
	}
	res = arr[0]
	for i := 0; i < len(arr); i++ {
		res ^= arr[i]
	}
	return res
}
