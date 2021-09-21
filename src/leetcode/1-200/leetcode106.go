package leetcode

func buildTree2(inorder []int, postorder []int) *TreeNode {
	var helper func(inorder []int, postorder []int, iStart, iEnd int, pStart, pEnd int) *TreeNode
	helper = func(inorder []int, postorder []int, iStart, iEnd int, pStart, pEnd int) *TreeNode {
		if iStart > iEnd || pStart > pEnd {
			return nil
		}
		index := 0
		for postorder[pEnd] != inorder[iStart+index] {
			index++
		}
		node := &TreeNode{Val: postorder[pEnd]}
		node.Left = helper(inorder, postorder, iStart, iStart+index-1, pStart, pStart+index-1)
		node.Right = helper(inorder, postorder, iStart+index+1, iEnd, pStart+index, pEnd-1)
		return node
	}
	return helper(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}
