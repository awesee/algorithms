package other

/*
Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

|.|8|7|6|5|4|3|2|1|
|2|.|.|.|.|.|.|.|.|
|3|.|.|.|.|.|.|.|.|
|4|.|.|.|.|.|.|.|.|
|5|.|.|.|.|.|.|.|.|
|6|.|.|.|.|.|.|.|.|
|7|.|.|.|.|.|.|.|.|
|8|.|.|.|.|.|.|.|.|
|9|.|.|.|.|.|.|.|.|

A partially filled sudoku which is valid.

Note:
A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.
*/

func IsValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {

		row := make(map[byte]bool)
		col := make(map[byte]bool)
		boxes := make(map[byte]bool)

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}

			if board[j][i] != '.' {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}

			if board[i/3*3+j/3][i%3*3+j%3] != '.' {
				if boxes[board[i/3*3+j/3][i%3*3+j%3]] {
					return false
				}
				boxes[board[i/3*3+j/3][i%3*3+j%3]] = true
			}
		}

	}

	return true
}
