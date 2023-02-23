package main

func main() {

}

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[byte]bool)
	colMap := make(map[int]map[byte]bool)
	boxMap := make(map[int]map[byte]bool)
	for i := 0; i < 9; i++ {
		rowMap[i] = map[byte]bool{}
		colMap[i] = map[byte]bool{}
		boxMap[i] = map[byte]bool{}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			box := (i/3)*3 + j/3

			if exit, ok := rowMap[i][board[i][j]]; ok && exit {
				return false
			} else {
				rowMap[i][board[i][j]] = true
			}

			if exit, ok := colMap[j][board[i][j]]; ok && exit {
				return false
			} else {
				colMap[j][board[i][j]] = true
			}

			if exit, ok := boxMap[box][board[i][j]]; ok && exit {
				return false
			} else {
				boxMap[box][board[i][j]] = true
			}
		}
	}
	return true
}
