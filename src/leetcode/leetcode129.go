package leetcode

func sumNumbers(root *TreeNode) int {
	res := 0
	var helper func(root *TreeNode, sum int)
	helper = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			res += sum
		}
		helper(root.Left, sum)
		helper(root.Right, sum)
	}
	helper(root, 0)
	return res
}
