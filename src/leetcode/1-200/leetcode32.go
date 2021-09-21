package leetcode

func longestValidParentheses(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if res < i-stack[len(stack)-1] {
					res = i - stack[len(stack)-1]
				}
			}
		}
	}
	return res
}
