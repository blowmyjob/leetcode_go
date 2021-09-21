package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return dfsMergeTwoLists(lists[0], lists[1])
	}
	n := len(lists)
	left := lists[0 : n/2]
	right := lists[n/2+1 : n]
	return dfsMergeTwoLists(mergeKLists(left), mergeKLists(right))
}

func dfsMergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	node := new(ListNode)
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		node.Val = l1.Val
		node.Next = dfsMergeTwoLists(l1.Next, l2)
	} else {
		node.Val = l2.Val
		node.Next = dfsMergeTwoLists(l1, l2.Next)
	}
	return node
}
