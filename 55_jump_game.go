// Source: https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	var max int = 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}