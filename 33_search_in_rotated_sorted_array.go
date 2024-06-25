// Source: https://leetcode.com/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	var l, r int = 0, len(nums) - 1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}
		if nums[l] <= nums[m] {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}