package leetcode

func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{}
	res := []int{}
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			tmp := queue[0]
			queue = queue[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
			if i == size-1 {
				res = append(res, tmp.Val)
			}
		}
	}
	return res
}
