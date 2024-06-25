// Source: https://leetcode.com/problems/sudoku-solver/
// Reference: https://www.geeksforgeeks.org/backtracking-algorithms/
func solveSudoku(board [][]byte)  {
	solve(board)
}

func solve(board [][]byte) bool {

for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					if isValid(board, i, j, byte(c)) {
						board[i][j] = byte(c)
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}


func isValid(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == c {
			return false
		}
		if board[i][col] == c {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == c {
			return false
		}
	}
	return true
}