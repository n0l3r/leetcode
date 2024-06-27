// Source: https://leetcode.com/problems/n-queens-ii/
func totalNQueens(n int) int {
	var res int
	var col, dia1, dia2 []bool
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			if !col[i] && !dia1[row+i] && !dia2[row-i+n-1] {
				col[i], dia1[row+i], dia2[row-i+n-1] = true, true, true
				dfs(row + 1)
				col[i], dia1[row+i], dia2[row-i+n-1] = false, false, false
			}
		}
	}
	col = make([]bool, n)
	dia1 = make([]bool, 2*n-1)
	dia2 = make([]bool, 2*n-1)
	dfs(0)
	return res
}