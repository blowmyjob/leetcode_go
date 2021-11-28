package contest

import (
	"sort"
)

func WateringPlants(plants []int, capacity int) int {
	size := 0
	water := capacity
	for i := 0; i < len(plants); i++ {
		if water < plants[i] {
			size += i * 2
			size++
			water = capacity - plants[i]
		} else {
			water -= plants[i]
			size++
		}
	}
	return size
}

type RangeFreqQuery struct{ pos [1e4 + 1]sort.IntSlice }

func Constructor1(arr []int) (q RangeFreqQuery) {
	for i, value := range arr {
		q.pos[value] = append(q.pos[value], i) // 统计 value 在 arr 中的所有下标位置
	}
	return
}

func (q *RangeFreqQuery) Query(left, right, value int) int {
	p := q.pos[value]                           // value 在 arr 中的所有下标位置
	return p[p.Search(left):].Search(right + 1) // 在下标位置上二分，求 [left,right] 之间的下标个数，即为 value 的频率
}
