package leetcode

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var resFunc func(root *TreeNode)
	resFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		resFunc(root.Left)
		resFunc(root.Right)
		res = append(res, root.Val)
	}
	resFunc(root)
	return res
}
