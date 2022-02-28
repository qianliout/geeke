package main

import (
	"strings"
)

func main() {

}

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
*/

func totalNQueens(n int) int {
	res := 0
	path := make([][]byte, 0)
	for i := 0; i < n; i++ {
		path = append(path, []byte(strings.Repeat(".", n)))
	}

	dfs(n, 0, path, &res)
	return res
}

func dfs(n, col int, path [][]byte, res *int) {
	if col == n {
		*res = (*res) + 1
		return
	}

	for row := 0; row < n; row++ {
		if valid(path, col, row, n) {
			path[col][row] = 'Q'
			dfs(n, col+1, path, res)
			path[col][row] = '.'
		}
	}
}
func valid(path [][]byte, col, row, n int) bool {
	// 横
	for i := 0; i < row; i++ {
		if path[col][i] == 'Q' {
			return false
		}
	}
	// 竖
	for i := 0; i < col; i++ {
		if path[i][row] == 'Q' {
			return false
		}
	}
	// 135度
	for i, j := col-1, row-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	// 45度
	for i, j := col-1, row+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if path[i][j] == 'Q' {
			return false
		}
	}

	return true
}
