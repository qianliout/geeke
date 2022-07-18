package main

import (
	"fmt"
)

func main() {
	// board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	// board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	// board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	// words := []string{"oath", "pea", "eat", "rain", "hklf", "hf"}
	board := [][]byte{{'a'}}
	// words := []string{"eat"}
	words := []string{"a"}
	ans := findWords(board, words)
	ans2 := findWords2(board, words)
	fmt.Println("ans is ", words, ans, ans2)
}

/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
*/
// 这种方式能得到解，但是会超时
func findWords2(board [][]byte, words []string) []string {
	ans := make([]string, 0)
	used := make(map[string]bool)

	for i := range words {
		tmp := make([]string, 0)

		for col := range board {
			for rol := range board[col] {
				if used[words[i]] {
					continue
				}
				dfs(board, col, rol, 0, []byte{}, []byte(words[i]), map[[2]int]bool{}, &tmp)
				if len(tmp) > 0 {
					ans = append(ans, words[i])
					used[words[i]] = true
				}
			}
		}

	}
	return ans
}
func dfs(board [][]byte, col, row, index int, path []byte, word []byte, used map[[2]int]bool, ans *[]string) {
	if index >= len(word) {
		*ans = append(*ans, string(word))
		return
	}
	if !in(board, col, row) || used[[2]int{col, row}] || board[col][row] != word[index] || len(*ans) > 0 {
		return
	}
	used[[2]int{col, row}] = true
	path = append(path, word[index])

	dfs(board, col+1, row, index+1, path, word, used, ans)
	dfs(board, col-1, row, index+1, path, word, used, ans)
	dfs(board, col, row+1, index+1, path, word, used, ans)
	dfs(board, col, row-1, index+1, path, word, used, ans)
	used[[2]int{col, row}] = false
	path = path[:len(path)-1]
}

func in(board [][]byte, col, row int) bool {
	if col < 0 || row < 0 || col >= len(board) || row >= len(board[col]) {
		return false
	}
	return true
}
