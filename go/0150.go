package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	top := -1
	for _, t := range tokens {
		switch t {
		case "+":
			stack[top-1] += stack[top]
			top--
		case "-":
			stack[top-1] -= stack[top]
			top--
		case "*":
			stack[top-1] *= stack[top]
			top--
		case "/":
			stack[top-1] /= stack[top]
			top--
		default:
			top++
			stack[top], _ = strconv.Atoi(t)
		}
	}
	return stack[0]
}
