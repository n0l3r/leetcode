// Source: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	var l, r int = 0, len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			i, j := m, m
			for i >= 0 && nums[i] == target {
				i--
			}
			for j < len(nums) && nums[j] == target {
				j++
			}
			return []int{i + 1, j - 1}
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return []int{-1, -1}
}