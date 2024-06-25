// Source: https://leetcode.com/problems/trapping-rain-water/
func trap(height []int) int {
	var res int
	var l int = 0
	var r int = len(height) - 1
	var lMax int
	var rMax int
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lMax {
				lMax = height[l]
			} else {
				res += lMax - height[l]
			}
			l++
		} else {
			if height[r] >= rMax {
				rMax = height[r]
			} else {
				res += rMax - height[r]
			}
			r--
		}
	}
	return res
}