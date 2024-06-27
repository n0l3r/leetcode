// Source: https://leetcode.com/problems/n-queens/
func solveNQueens(n int) [][]string {
	var res [][]string
	var board [][]byte = make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			var temp []string
			for i := 0; i < n; i++ {
				temp = append(temp, string(board[i]))
			}
			res = append(res, temp)
			return
		}
		for col := 0; col < n; col++ {
			if isValid(board, row, col) {
				board[row][col] = 'Q'
				dfs(row + 1)
				board[row][col] = '.'
			}
		}
	}
	dfs(0)
	return res
}

func isValid(board [][]byte, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}