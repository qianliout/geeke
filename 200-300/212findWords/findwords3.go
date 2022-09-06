package main

import (
	"fmt"
)

type TrieNode struct {
	val  rune
	word string
	next map[rune]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	// 构造字典树
	root := &TrieNode{next: map[rune]*TrieNode{}}
	for _, w := range words {
		for i, b := range w {
			if _, ok := root.next[b]; !ok {
				root.next[b] = &TrieNode{val: b, next: map[rune]*TrieNode{}}
			}
			if i == len(w)-1 {
				root.next[b].word = w
			}
			root = root.next[b]
		}
	}

	res := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			DFS(board, &res, root, i, j, map[string]bool{})
		}
	}

	return res
}

func DFS(brd [][]byte, res *[]string, node *TrieNode, i, j int, visited map[string]bool) {
	k := fmt.Sprintf("%v_%v", i, j)

	if i < 0 || j < 0 || i >= len(brd) || j >= len(brd[0]) || visited[k] {
		return
	}
	nn, ok := node.next[rune(brd[i][j])]
	if !ok {
		return
	}

	if len(nn.word) > 0 {
		*res = append(*res, nn.word)
		nn.word = ""
	}

	visited[k] = true
	DFS(brd, res, nn, i+1, j, visited)
	DFS(brd, res, nn, i-1, j, visited)
	DFS(brd, res, nn, i, j+1, visited)
	DFS(brd, res, nn, i, j-1, visited)
	visited[k] = false
}
