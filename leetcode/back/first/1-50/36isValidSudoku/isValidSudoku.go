package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}}

	res := isValidSudoku(board)
	fmt.Println("res is ", res)

}

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
    数字 1-9 在每一行只能出现一次。
    数字 1-9 在每一列只能出现一次。
    数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
func isValidSudoku(board [][]byte) bool {
	colMap := make(map[int]map[byte]bool)
	rowMap := make(map[int]map[byte]bool)
	miniMap := make(map[int]map[int]map[byte]bool)

	// 初始化各个map
	for i := 0; i < 9; i++ {
		colMap[i] = make(map[byte]bool)
		rowMap[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		miniMap[i] = make(map[int]map[byte]bool)
		for j := 0; j < 9; j++ {
			miniMap[i][j] = make(map[byte]bool)
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			// 先判断col

			if colMap[i][board[i][j]] {
				return false
			} else {
				colMap[i][board[i][j]] = true
			}
			// 再判断row
			if rowMap[j][board[i][j]] {
				return false
			} else {
				rowMap[j][board[i][j]] = true
			}

			// 再判断小方格
			col, row := i/3, j/3

			if miniMap[col][row][board[i][j]] {
				return false
			} else {
				miniMap[col][row][board[i][j]] = true
			}
		}
	}
	return true
}
