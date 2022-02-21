package main

func main() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	solveSudoku(board)
}

func solveSudoku(board [][]byte) {
	var hen, shu [9][9]bool
	var min [3][3][9]bool

	var spaces [][2]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				hen[i][board[i][j]-'1'] = true
				shu[j][board[i][j]-'1'] = true
				min[i/3][j/3][board[i][j]-'1'] = true
			}
		}
	}

	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx < 0 || idx >= len(spaces) {
			return true
		}

		i, j := spaces[idx][0], spaces[idx][1]
		for b := '1'; b <= '9'; b++ {
			dig := byte(b) - '1'

			if !hen[i][dig] && !shu[j][dig] && !min[i/3][j/3][dig] {
				board[i][j] = byte(b)
				hen[i][dig] = true
				shu[j][dig] = true
				min[i/3][j/3][dig] = true
				if dfs(idx + 1) {
					return true
				}
				board[i][j] = '.'
				hen[i][dig] = false
				shu[j][dig] = false
				min[i/3][j/3][dig] = false
			}
		}
		return false
	}
	if len(spaces) > 0 {
		dfs(0)
	}
}

func isValidSudoku(b byte, i, j int, hen, shu, min map[int]map[byte]bool) bool {
	if hen[i*10+j][b] {
		return false
	}
	if shu[i+j*10][b] {
		return false
	}
	if min[i/3*10+j/3][b] {
		return false
	}
	return true
}
