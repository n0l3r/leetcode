// Source: https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	var i int = 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}