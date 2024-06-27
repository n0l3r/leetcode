// Source: https://leetcode.com/problems/jump-game-ii/
func jump(nums []int) int {
	var res, end, max int
	for i := 0; i < len(nums)-1; i++ {
		if max < i+nums[i] {
			max = i + nums[i]
		}
		if i == end {
			end = max
			res++
		}
	}
	return res
}