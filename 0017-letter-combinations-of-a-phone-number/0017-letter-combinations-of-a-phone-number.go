// Source: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	var dict map[byte]string = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var dfs func(int, string)
	dfs = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		for _, c := range dict[digits[index]] {
			dfs(index+1, path+string(c))
		}
	}
	dfs(0, "")
	return res
}