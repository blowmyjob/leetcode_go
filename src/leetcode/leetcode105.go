package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	var helper func(preorder []int, inorder []int, pStart, pEnd int, iStart, iEnd int) *TreeNode
	helper = func(preorder []int, inorder []int, pStart, pEnd int, iStart, iEnd int) *TreeNode {
		if pStart > pEnd || iStart > iEnd {
			return nil
		}
		index := 0
		for preorder[pStart] != inorder[iStart+index] {
			index++
		}
		node := &TreeNode{Val: preorder[pStart]}
		node.Left = helper(preorder, inorder, pStart+1, pStart+index, iStart, iStart+index-1)
		node.Right = helper(preorder, inorder, pStart+index+1, pEnd, iStart+index+1, iEnd)
		return node
	}
	return helper(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}
