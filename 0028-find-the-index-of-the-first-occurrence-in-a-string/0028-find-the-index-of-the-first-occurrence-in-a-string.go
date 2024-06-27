// Source: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			break
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}