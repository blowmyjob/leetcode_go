package leetcode

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxTree(dfs(root.Left), dfs(root.Right))
}

func maxTree(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
