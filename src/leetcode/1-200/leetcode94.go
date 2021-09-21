package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var resFunc func(node *TreeNode)
	resFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		resFunc(node.Left)
		res = append(res, node.Val)
		resFunc(node.Right)
	}
	resFunc(root)
	return res
}
