package main

import (
	"fmt"

	"outback/leetcode/common/trie"
)

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}
	res := exist(board, "ABCCED")
	fmt.Println("res is ", res)
}

/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
示例:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
链接：https://leetcode-cn.com/problems/word-search
*/

// 回漱加dfs的方法
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}
	res := make([]string, 0)
	path := make([]byte, 0)
	used := make(map[int]map[int]bool)
	// 初始化map
	for i := 0; i < len(board); i++ {
		used[i] = make(map[int]bool)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, path, []byte(word), i, j, 0, used, &res)
			if len(res) > 0 {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, path, word []byte, i, j, k int, used map[int]map[int]bool, res *[]string) {
	if len(path) == len(word) {
		*res = append(*res, string(append([]byte{}, path...)))
	}
	//fmt.Println("dfs res is ", *res)
	// 这里一定要注意，这两个判断要写在加res后面，因为当path符合条件时，可能i,j,k已达到最后不符全条件，因为在上一个
	// 循环中增加了一
	if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 || k < 0 || k >= len(word) {
		return
	}
	if len(*res) > 0 {
		return // 找到一个就返回
	}
	if board[i][j] == []byte(word)[k] {
		if !used[i][j] {
			path = append(path, board[i][j])
			k++
			used[i][j] = true
			dfs(board, path, word, i+1, j, k, used, res)
			dfs(board, path, word, i-1, j, k, used, res)
			dfs(board, path, word, i, j+1, k, used, res)
			dfs(board, path, word, i, j-1, k, used, res)
			path = path[:len(path)-1]
			k--
			used[i][j] = false
		}
	}
}

// 字典树的方法
func exist2(board [][]byte, word string) bool {
	return true
}

func newTrie(board [][]byte) *trie.Trie {
	t := trie.Constructor()
	return &t

}
