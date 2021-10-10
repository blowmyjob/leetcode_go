package leetcode

func rob2(root *TreeNode) int {
	var helper func(root *TreeNode) []int
	helper = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		result := []int{0, 0}
		left := helper(root.Left)
		right := helper(root.Right)
		result[0] = maxResult(left[0], left[1]) + maxResult(right[0], right[1])
		result[1] = root.Val + left[0] + right[0]
		return result
	}
	result := helper(root)
	return maxResult(result[0], result[1])
}

func maxResult(i, j int) int {
	if i > j {
		return i
	}
	return j
}
