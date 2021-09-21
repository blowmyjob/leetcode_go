package leetcode

func levelOrderBottom(root *TreeNode) [][]int {
	res1 := [][]int{}
	res := [][]int{}
	if root == nil {
		return res1
	}
	arr := []*TreeNode{}
	arr = append(arr, root)
	for len(arr) != 0 {
		size := len(arr)
		tmp := []int{}
		for i := 0; i < size; i++ {
			node := arr[0]
			tmp = append(tmp, node.Val)
			arr = arr[:size-1]
			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		res1 = append(res1, tmp)
	}
	for i := len(res1) - 1; i >= 0; i-- {
		res = append(res, res1[i])
	}
	return res
}
