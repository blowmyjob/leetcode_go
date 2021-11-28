package contest

import (
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool { return items[i][0] < items[j][0] }) // 按价格排序
	for i, q := range queries {
		queries[i] = q<<32 | i // 这样排序时可以保留查询的下标
	}
	sort.Ints(queries)

	ans := make([]int, len(queries))
	maxBeauty, i := 0, 0
	for _, q := range queries {
		for ; i < len(items) && items[i][0] <= q>>32; i++ {
			if items[i][1] > maxBeauty {
				maxBeauty = items[i][1]
			}
		}
		ans[q&(1<<32-1)] = maxBeauty
	}
	return ans
}

func ReverseEvenLengthGroups(head *ListNode) *ListNode {
	res := []int{0, 4, 2, 1, 3}

	left, right, groupNum := 0, 0, 1
	flag1, flag2 := true, false
	for right < len(res) && flag1 {
		if (right-left+1)%2 == 0 {
			for i, j := left, right; i <= j; i, j = i+1, j-1 {
				res[i], res[j] = res[j], res[i]
			}
			if flag2 {
				flag1 = false
			}
		}
		groupNum++
		left = right + 1
		if right+groupNum >= len(res) {
			right = len(res) - 1
			flag2 = true
		} else {
			right = right + groupNum
		}
	}
	p := &ListNode{}
	tmp := p
	for i := 0; i < len(res); i++ {
		tmp.Next = &ListNode{Val: res[i]}
		tmp = tmp.Next
	}
	return p.Next
}

func DecodeCiphertext(encodedText string, rows int) string {
	length := len(encodedText)
	cols := length / rows
	ans := make([]byte, 0)
	for index := 0; index < cols; index++ {
		for i, j := 0, index; i < rows && j < cols; i, j = i+1, j+1 {
			ans = append(ans, encodedText[i*cols+j])
		}
	}
	for i := len(ans) - 1; i >= 0 && ans[i] == byte(' '); i-- {
		ans = ans[:i]
	}
	return string(ans)
}
