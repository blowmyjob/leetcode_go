package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	res := [][]int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		tmpArr := []int{}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmpArr = append(tmpArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmpArr)
	}
	return res
}
