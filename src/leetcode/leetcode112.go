package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs112(root, targetSum)
}

func dfs112(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	sum -= root.Val
	return dfs112(root.Left, sum) || dfs112(root.Right, sum)
}
