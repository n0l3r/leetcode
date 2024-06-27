// Source: https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	var l, r int = 0, len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}