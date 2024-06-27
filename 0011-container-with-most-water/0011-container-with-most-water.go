// Source: https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	var res int
	var l int = 0
	var r int = len(height) - 1
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}