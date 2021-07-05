package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NumberOf1(n int) int {
	// write code here
	var count int = 0
	for int32(n) != 0 {
		n &= int(int32(n - 1))
		count++
	}
	return count
}

func reOrderArray(array []int) []int {
	// write code here
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j]%2 == 0 && array[j+1]%2 != 0 {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func FindKthToTail(pListHead *ListNode, k int) *ListNode {
	// write code here
	if pListHead == nil {
		return nil
	}
	pre := pListHead
	for i := 0; i < k; i++ {
		if pre == nil {
			return nil
		}
		pre = pre.Next
	}
	cur := pListHead
	for pre != nil {
		pre = pre.Next
		cur = cur.Next
	}
	return cur
}

func ReverseList(pHead *ListNode) *ListNode {
	var prev *ListNode
	for pHead != nil {
		prev, pHead, pHead.Next = pHead, pHead.Next, prev
	}
	return prev

	//if pHead == nil {
	//	return nil
	//}
	//
	//if pHead.Next == nil {
	//	return pHead
	//}
	//
	//a := ReverseList(pHead.Next)
	//pHead.Next.Next = pHead
	//pHead.Next = nil
	//
	//return a
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	if pHead1.Val < pHead2.Val {
		pHead1.Next = Merge(pHead1.Next, pHead2)
		return pHead1
	} else {
		pHead2.Next = Merge(pHead1, pHead2.Next)
		return pHead2
	}
}

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// write code here
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	return CheckTreeNode(pRoot1, pRoot2) || HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func CheckTreeNode(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot2 == nil {
		return true
	}
	if pRoot1 == nil {
		return false
	}
	return pRoot1.Val == pRoot2.Val && CheckTreeNode(pRoot1.Left, pRoot1.Left) && CheckTreeNode(pRoot1.Right, pRoot1.Right)
}

func Mirror(pRoot *TreeNode) *TreeNode {
	// write code here
	if pRoot == nil {
		return pRoot
	}
	pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}

var (
	stk = make([]int, 0)
	min = make([]int, 0)
)

func Push(node int) {
	// write code here
	if len(min) > 0 && min[len(min)-1] < node {
		min = append(min, min[len(min)-1])
	} else {
		min = append(min, node)
	}
	stk = append(stk, node)
}
func Pop() {
	// write code here
	stk = stk[:len(stk)-1]
	min = min[:len(min)-1]
}
func Top() int {
	// write code here
	return stk[len(stk)-1]
}
func Min() int {
	// write code here
	return min[len(min)-1]
}

func IsPopOrder(pushV []int, popV []int) bool {
	j := 0
	i := 0
	b := []int{}
	for i < len(popV) {
		if pushV[i] == popV[j] {
			j++
			i++
			for len(b) != 0 && b[len(b)-1] == popV[j] {
				b = b[:len(b)-1]
				j++
			}
		} else {
			b = append(b, pushV[i])
			i++
		}
	}
	return len(b) == 0
}

func PrintFromTopToBottom(root *TreeNode) []int {
	// write code here
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := make([]*TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			res = append(res, queue[i].Val)
			if queue[i].Left != nil {
				tmp = append(tmp, queue[i].Left)
			}
			if queue[i].Right != nil {
				tmp = append(tmp, queue[i].Right)
			}
		}
		queue = tmp
	}
	return res
}
