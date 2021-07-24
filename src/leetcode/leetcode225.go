package leetcode

type MyStack struct {
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		stack1: []int{},
		stack2: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.stack2 = append(this.stack2, x)
	for len(this.stack1) > 0 {
		tmp := this.stack1[0]
		this.stack1 = this.stack1[1:]
		this.stack2 = append(this.stack2, tmp)
	}
	for len(this.stack2) > 0 {
		tmp := this.stack2[0]
		this.stack2 = this.stack2[1:]
		this.stack1 = append(this.stack1, tmp)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.stack1[0]
	this.stack1 = this.stack1[1:]
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.stack1) == 0
}
