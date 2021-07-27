package leetcode

import "math"

func isValidBST(root *TreeNode) bool {
	return dfsBST(root, math.MinInt64, math.MaxInt64)
}

func dfsBST(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return dfsBST(root.Left, lower, root.Val) && dfsBST(root.Right, root.Val, upper)
}
