// Source: https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, []int)
	backtrack = func(nums []int, track []int) {
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if contains(track, nums[i]) {
				continue
			}
			track = append(track, nums[i])
			backtrack(nums, track)
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, []int{})
	return res  
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}