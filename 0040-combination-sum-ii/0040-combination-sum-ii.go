// Source: https://leetcode.com/problems/combination-sum-ii/
func combinationSum2(candidates []int, target int) [][]int {
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
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if target < candidates[i] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	sort.Ints(candidates)
	backtrack(0, target, []int{})
	return res
}