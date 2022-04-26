package main

func main() {

}

/*
给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
*/
func solve(board [][]byte) {
	// 先找出四条边相连的所有'0'
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		dfs(board, i, 0, 'R', 'O')
		dfs(board, i, len(board[0])-1, 'R', 'O')

	}
	for i := 0; i < len(board[0]); i++ {
		dfs(board, 0, i, 'R', 'O')
		dfs(board, len(board)-1, i, 'R', 'O')
	}
	// 估后把里面的所有的换成'X'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	// 把其他的换回来
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, col, row int, fill, pre byte) {
	if !in(board, col, row) {
		return
	}
	if board[col][row] != pre {
		return
	}
	if board[col][row] == pre {
		board[col][row] = fill
	}
	dfs(board, col+1, row, fill, pre)
	dfs(board, col-1, row, fill, pre)
	dfs(board, col, row+1, fill, pre)
	dfs(board, col, row-1, fill, pre)
}

func in(board [][]byte, col, row int) bool {
	if col < 0 || col >= len(board) || row < 0 || row >= len(board[0]) {
		return false
	}
	return true
}
