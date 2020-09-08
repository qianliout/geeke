package main

import (
	"fmt"
)

func main() {
	res := solveNQueens(4)
	fmt.Println(len(res))
}

/*
 n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
上图为 8 皇后问题的一种解法。
给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
示例：
输入：4
输出：[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。
提示：
    皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
*/

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([][]byte, 0)
	for i := 0; i < n; i++ {
		r := make([]byte, 0)
		for j := 0; j < n; j++ {
			r = append(r, '.')
		}
		path = append(path, r)
	}

	dfs(0, path, &res, n)
	return res

}

func dfs(col int, path [][]byte, res *[][]string, n int) {
	if col >= n {
		recs := make([]string, 0)
		for _, s := range path {
			recs = append(recs, string(s))
		}
		*res = append(*res, append([]string{}, recs...))
		return
	}
	for row := 0; row < n; row++ {
		if isValid(col, row, path, n) {
			path[col][row] = 'Q'
			dfs(col+1, path, res, n)
			path[col][row] = '.'
		}
	}
}

// 检查在原来已填充了path的情况下，再填col,row这个点，会不会有问题，检查的原则就是，这几个规则下原来还没有过Q
func isValid(col, row int, path [][]byte, n int) bool {
	// 检查列
	for i := 0; i < col; i++ {
		if path[i][row] == 'Q' {
			return false
		}
	}
	// 再检查行
	for i := 0; i < row; i++ {
		if path[col][i] == 'Q' {
			return false
		}
	}
	// 再检查135度
	for i, j := col-1, row-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	// 再检查45度
	for i, j := col-1, row+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if path[i][j] == 'Q' {
			return false
		}
	}
	return true
}
