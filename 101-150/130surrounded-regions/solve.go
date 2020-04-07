package main

import (
	"fmt"

	"outback/leetcode/common/unionfind"
)

func main() {
	board := [][]byte{
		[]byte("XOXOXO"),
		[]byte("OXOXOX"),
		[]byte("XOXOXO"),
		[]byte("OXOXOX"),
	}

	solve(board)
	solve2(board)
}

/*
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
示例:
X X X X
X O O X
X X O X
X O X X
运行你的函数后，矩阵变为：
X X X X
X X X X
X X X X
X O X X
解释:
被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。
任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
链接：https://leetcode-cn.com/problems/surrounded-regions
*/

func solve(board [][]byte) {
	if len(board) <= 1 || len(board[0]) <= 1 {
		return
	}
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			isEnd := i == 0 || i == m-1 || j == 0 || j == n-1

			if isEnd && board[i][j] == 'O' {
				DFS(&board, i, j)
			}
		}
	}
	for _, nums := range board {
		fmt.Println(string(nums))
	}
	// 换值
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
	for _, nums := range board {
		fmt.Println(string(nums))
	}
	//fmt.Println("res is ", board)

}

func DFS(board *[][]byte, col, row int) {
	if row < 0 || row >= len((*board)[0]) || col < 0 || col >= len(*board) || (*board)[col][row] == 'X' || (*board)[col][row] == '#' {
		return
	}
	(*board)[col][row] = '#'
	DFS(board, col+1, row)
	DFS(board, col-1, row)
	DFS(board, col, row+1)
	DFS(board, col, row-1)
}

// 本题也可以使用迭代的方式，便是就没有递归代码容易理解

func solve2(board [][]byte) {
	if len(board) <= 1 || len(board[0]) <= 1 {
		return
	}
	m := len(board)
	n := len(board[0])
	uf := unionfind.NewUnionFind(m*n + 1) // 这里加一的目的是增加一个dump node
	dummyNode := m * n
	fmt.Println(uf)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					uf.Union(Node(i, j, n), dummyNode)
				} else {
					if i > 0 && board[i-1][j] == 'O' {
						uf.Union(Node(i, j, n), Node(i-1, j, n))
					}
					if i < m-1 && board[i+1][j] == 'O' {
						uf.Union(Node(i, j, n), Node(i+1, j, n))
					}

					if j > 0 && board[i][j-1] == 'O' {
						uf.Union(Node(i, j, n), Node(i, j-1, n))
					}
					if j < n-1 && board[i][j+1] == 'O' {
						uf.Union(Node(i, j, n), Node(i, j+1, n))
					}
				}
			}
		}
	}
	//fmt.Println(uf.Parent)
	//fmt.Println(uf.Rank)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if uf.IsConnected(Node(i, j, n), dummyNode) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
	for _, nums := range board {
		fmt.Println(string(nums))
	}
}

func Node(x, y, length int) int {
	return x*length + y
}
