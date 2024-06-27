// Source: https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backtrack func(int, int, []int)
	backtrack = func(start, target int, path []int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if target < candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return res
}