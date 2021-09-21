package leetcode

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var resFunc func(root *TreeNode)
	resFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		resFunc(root.Left)
		resFunc(root.Right)
	}
	resFunc(root)
	return res
}
