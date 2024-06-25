// Source: https://leetcode.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var dfs func(int)
	sort.Ints(nums)
	dfs = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			res = append(res, temp)
			return
		}
		m := make(map[int]bool)
		for i := index; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			m[nums[i]] = true
			nums[index], nums[i] = nums[i], nums[index]
			dfs(index + 1)
			nums[index], nums[i] = nums[i], nums[index]
		}
	}
	dfs(0)
	return res
}