// Source: https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	var res int
	var sign int = 1
	var i int
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		if res*sign > 1<<31-1 {
			return 1<<31 - 1
		}
		if res*sign < -(1 << 31) {
			return -(1 << 31)
		}
		i++
	}
	return res * sign
}