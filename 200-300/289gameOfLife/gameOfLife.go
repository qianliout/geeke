package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	bored := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(bored)
	fmt.Println(bored)

}

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	for col := range board {
		for row := range board[col] {
			live := find(board, col, row)
			// 原来是活的，现在死了
			if (live < 2 || live > 3) && board[col][row] == 1 {
				board[col][row] = -1
			}
			// 原来是死的，现在活了
			if live == 3 && board[col][row] == 0 {
				board[col][row] = 2
			}
		}
	}

	for col := range board {
		for row := range board[col] {
			if board[col][row] <= 0 {
				board[col][row] = 0
			}
			if board[col][row] >= 1 {
				board[col][row] = 1
			}
		}
	}
}

func find(board [][]int, col, row int) int {
	ans := 0
	for i := col - 1; i <= col+1; i++ {
		for j := row - 1; j <= row+1; j++ {
			if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || (i == col && j == row) {
				continue
			}
			if Abs(board[i][j]) == 1 {
				ans++
			}
		}
	}
	return ans
}
