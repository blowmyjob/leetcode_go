package leetcode

func flatten(root *TreeNode) {
	list := transfer(root)
	for i := 1; i < len(list); i++ {
		pre, curr := list[i-1], list[i]
		pre.Left, pre.Right = nil, curr
	}
}

func transfer(root *TreeNode) []*TreeNode {
	list := make([]*TreeNode, 0)
	if root != nil {
		list = append(list, root)
		list = append(list, transfer(root.Left)...)
		list = append(list, transfer(root.Right)...)
	}
	return list
}
