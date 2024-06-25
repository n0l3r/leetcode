// Source: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var maxLen int
	var start int
	var charMap = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if charMap[s[i]] > 0 {
			start = max(start, charMap[s[i]])
		}
		charMap[s[i]] = i + 1
		maxLen = max(maxLen, i-start+1)
	}
	return maxLen
}