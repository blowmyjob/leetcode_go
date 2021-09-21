package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	res := []int{}
	var orderFunc func(root *Node)
	orderFunc = func(root *Node) {
		if root == nil {
			return
		}
		for _, v := range root.Children {
			orderFunc(v)
		}
		res = append(res, root.Val)
	}
	orderFunc(root)
	return res
}
