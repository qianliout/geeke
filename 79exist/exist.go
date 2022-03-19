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
		if word == "" {
			return true
		}
		return false
	}

	dp := make([][]bool, len(board))
	for i := range dp {
		dp[i] = make([]bool, len(board[0]))
	}

	return false
}

func dfs() {

}
