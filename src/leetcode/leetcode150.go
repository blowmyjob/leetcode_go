package leetcode

import "strconv"

func EvalRPN(tokens []string) int {
	stack := []string{}
	res := 0
	for i := 0; i < len(tokens); i++ {
		_, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, tokens[i])
		}
		if tokens[i] == "+" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			num := num1 + num2
			stack = append(stack, strconv.Itoa(num))
		}
		if tokens[i] == "*" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			num := num1 * num2
			stack = append(stack, strconv.Itoa(num))
		}
		if tokens[i] == "-" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			num := num1 - num2
			stack = append(stack, strconv.Itoa(num))
		}
		if tokens[i] == "/" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]
			num := num1 / num2
			stack = append(stack, strconv.Itoa(num))
		}
	}
	res, _ = strconv.Atoi(stack[0])
	return res
}
