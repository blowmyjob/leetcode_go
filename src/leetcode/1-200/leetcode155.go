package leetcode

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		stack:    []int{},
//		minStack: []int{math.MaxInt64},
//	}
//}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, stackMin(top, val))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func stackMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
