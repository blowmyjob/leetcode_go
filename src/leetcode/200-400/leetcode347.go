package leetcode

import "container/heap"

type Iheap [][2]int

func topKFrequent(nums []int, k int) []int {
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := maps[nums[i]]; ok {
			maps[nums[i]]++
		} else {
			maps[nums[i]] = 1
		}
	}
	iheap := &Iheap{}
	heap.Init(iheap)
	for key, value := range maps {
		if iheap.Len() < k {
			heap.Push(iheap, [2]int{key, value})
		} else {
			if (*iheap)[0][1] < value {
				heap.Pop(iheap)
				heap.Push(iheap, [2]int{key, value})
			}
		}
	}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(iheap).([2]int)[0]
	}

	return res
}

func (h Iheap) Len() int           { return len(h) }
func (h Iheap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h Iheap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Iheap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *Iheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
