// Source: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}