package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := new(ListNode)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		node.Val = l1.Val
		node.Next = mergeTwoLists(l1.Next, l2)
	} else {
		node.Val = l2.Val
		node.Next = mergeTwoLists(l1, l2.Next)
	}
	return node
}
