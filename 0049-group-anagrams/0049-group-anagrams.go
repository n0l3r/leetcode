// Source: https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	var m = make(map[string][]string)
	for _, str := range strs {
		var b = []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		m[string(b)] = append(m[string(b)], str)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}