package trie

import "fmt"

type Trienode struct {
	val  rune
	word string
	next map[rune]*Trienode
}

func FindWords(board [][]byte, words []string) []string {
	// 构造字典树
	root := &Trienode{next: map[rune]*Trienode{}}
	for _, w := range words {
		p := root
		for i, b := range w {
			if _, ok := p.next[b]; !ok {
				p.next[b] = &Trienode{val: b, next: map[rune]*Trienode{}}
			}
			if i == len(w)-1 {
				p.next[b].word = w
			}
			p = p.next[b]
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

func DFS(brd [][]byte, res *[]string, node *Trienode, i, j int, visited map[string]bool) {
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
