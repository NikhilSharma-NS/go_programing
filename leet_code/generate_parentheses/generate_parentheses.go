package generate_parentheses

import (
	"strings"
)

func GenerateParenthesis(n int) []string {
	stack := []string{}
	res := []string{}
	backtrack(0, 0, n, stack, &res)
	return res
}

func backtrack(open int, close int, n int, stack []string, result *[]string) {
	if open == close && close == n && open == n {
		s := strings.Join(stack, "")
		*result = append(*result, s)
		return
	}
	if open < n {
		stack = append(stack, "(")
		backtrack(open+1, close, n, stack, result)
		stack = stack[:len(stack)-1]
	}
	if close < open {
		stack = append(stack, ")")
		backtrack(open, close+1, n, stack, result)
		stack = stack[:len(stack)-1]
	}
}
