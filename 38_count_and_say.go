// Source: https://leetcode.com/problems/count-and-say/
func countAndSay(n int) string {
	var res = "1"
	for i := 1; i < n; i++ {
		var temp string
		var count int
		for j := 0; j < len(res); j++ {
			count++
			if j == len(res)-1 || res[j] != res[j+1] {
				temp += strconv.Itoa(count) + string(res[j])
				count = 0
			}
		}
		res = temp
	}
	return res
}