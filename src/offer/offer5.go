package offer

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		stack2 = append(stack2, stack1...)
		stack1 = []int{}
	}
	rs := stack2[0]
	stack2 = stack2[1:]
	return rs
}
