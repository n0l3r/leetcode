// Source: https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int)  {
	var i int = len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		var j int = len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)

	
}

func reverse(nums []int, start int) {
	var i int = start
	var j int = len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}