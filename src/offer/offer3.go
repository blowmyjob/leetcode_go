package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	res := make([]int, 0)
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}
	return res
}
