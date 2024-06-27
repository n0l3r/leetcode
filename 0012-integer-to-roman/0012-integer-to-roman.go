// Source: https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	var res string
	var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			res += symbols[i]
		}
	}
	return res
}