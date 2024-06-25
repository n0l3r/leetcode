// Source: https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	var res, sum int = nums[0], 0
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > res {
			res = sum
		}
	}
	return res
}