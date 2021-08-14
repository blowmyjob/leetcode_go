package leetcode

func oddEvenList(head *ListNode) *ListNode {
	oddList := &ListNode{}
	oddHead := oddList
	evenList := &ListNode{}
	evenHead := evenList
	p := head
	isOdd := true
	for p != nil {
		if isOdd {
			oddList.Next = p
			oddList = oddList.Next
		} else {
			evenList.Next = p
			evenList = evenList.Next
		}
		isOdd = !isOdd
		p = p.Next
	}
	oddList.Next = evenHead.Next
	evenList.Next = nil
	return oddHead.Next
}
