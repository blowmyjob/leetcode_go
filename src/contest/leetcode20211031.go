package contest

import (
	"math"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	p := head
	res := []int{}
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	arr := []int{}
	for i := 1; i < len(res)-1; i++ {
		if (res[i-1] < res[i] && res[i] > res[i+1]) || (res[i-1] > res[i] && res[i] < res[i+1]) {
			arr = append(arr, i)
		}
	}
	sort.Ints(arr)
	min := math.MaxInt32
	for i := 0; i < len(arr)-1; i++ {
		if min > int(math.Abs(float64(arr[i+1]-arr[i]))) {
			min = int(math.Abs(float64(arr[i+1] - arr[i])))
		}
	}
	if min == math.MaxInt32 || len(arr) <= 0 {
		return []int{-1, -1}
	}
	max := arr[len(arr)-1] - arr[0]
	return []int{min, max}
}

func smallestEqual(nums []int) int {
	wordMap := make(map[int]int)
	for k, v := range nums {
		wordMap[k] = v
	}
	for i := 0; i < len(nums); i++ {
		tmp := i % 10
		if nums[i] == tmp {
			return i
		}
	}
	return -1
}

func minimumOperations(nums []int, start int, goal int) int {
	N := 1000
	dst := make([]int, N+1)
	for i := range dst {
		dst[i] = math.MaxInt32
	}
	dst[start] = 0
	queue := []int{start}
	var v int
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		for _, num := range nums {
			for i := 0; i < 3; i++ {
				if i == 0 {
					v = x + num
				} else if i == 1 {
					v = x - num
				} else {
					v = x ^ num
				}
				if v == goal {
					return dst[x] + 1
				}
				if v < 0 || v > N {
					continue
				}
				if dst[x]+1 < dst[v] {
					dst[v] = dst[x] + 1
					queue = append(queue, v)
				}
			}
		}
	}
	return -1
}
