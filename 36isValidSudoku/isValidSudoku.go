package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	sudoku := isValidSudoku(board)
	fmt.Println("sudoku", sudoku)
}

/*
请你判断一个 9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。（请参考示例图）
*/
func isValidSudoku(board [][]byte) bool {
	// hen, shu, quan := make(map[int]byte), make(map[int]byte), make(map[int]byte)

	for i := range board {
		hen := make(map[byte]bool)
		for j := range board[i] {
			if board[i][j] != '.' && hen[board[i][j]] {
				return false
			}
			hen[board[i][j]] = true
		}
	}
	for i := 0; i < 9; i++ {
		shu := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' && shu[board[j][i]] {
				return false
			}
			shu[board[j][i]] = true
		}
	}

	miniMap := make(map[string]map[byte]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			miniMap[fmt.Sprintf("%d%d", i, j)] = make(map[byte]bool)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// 再判断小方格
			col, row := i/3, j/3
			key := fmt.Sprintf("%d%d", col, row)

			if board[i][j] != '.' && miniMap[key][board[i][j]] {
				return false
			}
			miniMap[key][board[i][j]] = true
		}
	}

	return true
}
