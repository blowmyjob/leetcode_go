package leetcode

type BSTIterator struct {
	arr []int
}

//func Constructor(root *TreeNode) (it BSTIterator) {
//	it.inOrder(root)
//	return
//}

func (this *BSTIterator) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.arr) > 0
}

func (this *BSTIterator) inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	this.inOrder(root.Left)
	this.arr = append(this.arr, root.Val)
	this.inOrder(root.Right)
}
