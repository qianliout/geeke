package main

import (
	"fmt"
	"strings"
)

func main() {
	queens := solveNQueens(4)
	fmt.Println("queues is ", queens)
}

/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
任意两个皇后不能位于同一行，同一列，同一斜线
*/
func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([][]byte, 0)
	for i := 0; i < n; i++ {
		path = append(path, []byte(strings.Repeat(".", n)))
	}

	dfs(n, 0, path, &res)
	return res
}

func dfs(n, col int, path [][]byte, res *[][]string) {
	if col == n {
		th := make([]string, 0)
		for i := range path {
			th = append(th, string(path[i]))
		}
		*res = append(*res, append([]string{}, th...))
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
