package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

func addTwo(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	if l1 != nil {
		carry += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		carry += l2.Val
		l2 = l2.Next
	}
	return &ListNode{
		Val:  carry % 10,
		Next: addTwo(l1, l2, carry/10),
	}
}
