func calculate(s string) int {
	res, num, sign := 0, 0, 1
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num * 10 + int(s[i] - '0')
		} else if s[i] == '+' {
			res += sign * num
			num = 0
			sign = 1
		} else if s[i] == '-' {
			res += sign * num
			num = 0
			sign = -1
		} else if s[i] == '(' {
			stack = append(stack, res)
			stack = append(stack, sign)
			res = 0
			sign = 1
		} else if s[i] == ')' {
			res += sign * num
			num = 0
			res *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	res += sign * num
	return res
}