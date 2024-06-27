// Source: https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			l, r := i+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[l] + nums[r]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < 0 {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}