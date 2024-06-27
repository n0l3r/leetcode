// Source: https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var res int
	var temp int = x
	for temp != 0 {
		res = res*10 + temp%10
		temp = temp / 10
	}
	return res == x
}