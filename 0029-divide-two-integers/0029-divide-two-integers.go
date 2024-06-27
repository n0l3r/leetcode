// Source: https://leetcode.com/problems/divide-two-integers/
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == -(1<<31) && divisor == -1 {
		return 1<<31 - 1
	}
	var sign int
	if dividend < 0 {
		sign--
		dividend = -dividend
	}
	if divisor < 0 {
		sign--
		divisor = -divisor
	}
	var res int
	for dividend >= divisor {
		var temp int = divisor
		var multiple int = 1
		for dividend >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}
		res += multiple
		dividend -= temp
	}
	if sign == -1 {
		return -res
	}
	return res
}