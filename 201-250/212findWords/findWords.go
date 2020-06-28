package main

import (
	"fmt"

	. "outback/leetcode/common/trie"
)

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'a', 'a'},
		//{'i', 'h', 'k', 'r'},
		//{'i', 'f', 'l', 'v'},
	}
	//words := []string{"oath", "pea", "eat", "rain"}
	words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	res := findWords(board, words)
	fmt.Println("res is ", res)
}

/*
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母在一个单词中不允许被重复使用。
示例:
输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]
["eat","oath"]
*/
func findWords(board [][]byte, words []string) []string {
	result := make([]string, 0)
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return result
	}

	tr := Constructor()
	for _, word := range words {
		tr.Insert(word)
	}
	used := make(map[string]bool)
	res := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, &res, "", &tr, j, i, used)
		}
	}
	//注意res这里去不了重，是为什么呢
	set := make(map[string]interface{})
	for _, i2 := range res {
		set[i2] = nil
	}

	for i := range set {
		result = append(result, i)
	}
	//fmt.Println(res)
	return result
}

func dfs(board [][]byte, res *[]string, path string, trieNode *Trie, row, col int, used map[string]bool) {
	v := fmt.Sprintf("%d_%d", col, row)
	if trieNode.Search(path) {
		*res = append(*res, path)
		path = ""
		return
	}
	if row < 0 || col < 0 || row >= len(board[0]) || col >= len(board) || used[v] {
		return
	}

	used[v] = true
	path = path + string(board[col][row])

	dfs(board, res, path, trieNode, row+1, col, used)
	dfs(board, res, path, trieNode, row-1, col, used)
	dfs(board, res, path, trieNode, row, col+1, used)
	dfs(board, res, path, trieNode, row, col-1, used)
	used[v] = false
	path = string([]byte(path)[:len(path)-1])
}
