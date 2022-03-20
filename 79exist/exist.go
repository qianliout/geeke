package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("sha256:00014a801a96a80d7ad68fb5bfb76dd374942d634b17b66cffc041daa4467c8f"))
}

/*
给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return word == ""
	}

	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[i]))
	}
	path, wordByte := make([]byte, 0), []byte(word)
	var res []string
	for i := range board {
		for j := range board[i] {
			dfs(board, i, j, 0, path, wordByte, &res, used)
			if len(res) > 0 {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, col, row, wordStart int, path, word []byte, res *[]string, used [][]bool) {
	if string(path) == string(word) {
		*res = append(*res, string(append([]byte{}, path...)))
		return
	}
	if len(*res) > 0 {
		return
	}
	if col < 0 || row < 0 || col >= len(board) || row >= len(board[0]) || wordStart < 0 || wordStart >= len(word) {
		return
	}
	if board[col][row] == word[wordStart] && !used[col][row] {
		path = append(path, board[col][row])
		used[col][row] = true
		wordStart++
		dfs(board, col+1, row, wordStart, path, word, res, used)
		dfs(board, col, row+1, wordStart, path, word, res, used)
		dfs(board, col-1, row, wordStart, path, word, res, used)
		dfs(board, col, row-1, wordStart, path, word, res, used)
		path = path[:len(path)-1]
		used[col][row] = false
		wordStart--
	}
}
