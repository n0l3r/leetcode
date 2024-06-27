// Source: https://leetcode.com/problems/longest-valid-parentheses/
func longestValidParentheses(s string) int {
	var res int
	var stack []int
	stack = append(stack, -1)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}