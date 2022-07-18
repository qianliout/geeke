package main

type Node struct {
	End  bool
	Next map[byte]*Node
}

func NewNode() *Node {
	return &Node{
		End:  false,
		Next: map[byte]*Node{},
	}
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: NewNode()}
}

func (t *Trie) Find(word string) bool {
	return find(t.Root, []byte(word), 0)
}
func (t *Trie) Add(word string) {
	add(t.Root, []byte(word), 0)
}

func add(node *Node, word []byte, index int) {
	if index >= len(word) {
		node.End = true
		return
	}
	c := word[index]
	if node.Next[c] == nil {
		node.Next[c] = NewNode()
	}
	add(node.Next[c], word, index+1)
}

func find(node *Node, word []byte, index int) bool {
	if index >= len(word) {
		return node.End
	}
	c := word[index]
	if node.Next[c] == nil {
		return false
	}
	return find(node.Next[c], word, index+1)
}

func findWords3(board [][]byte, words []string) []string {
	trie := Constructor()
	for i := range words {
		trie.Add(words[i])
	}
	ans := make([]string, 0)
	res := make(map[string]bool)
	// 找值
	for col := range board {
		for row := range board[col] {
			DFs(board, res, trie.Root, col, row, map[[2]int]bool{}, []byte{})
		}
	}
	for k := range res {
		ans = append(ans, k)
	}

	return ans
}

func DFs(brd [][]byte, res map[string]bool, node *Node, col, row int, visited map[[2]int]bool, path []byte) {
	if node.End {
		res[string(path)] = true
		return
	}
	if col < 0 || row < 0 || col >= len(brd) || row >= len(brd[0]) || visited[[2]int{col, row}] {
		return
	}

	c := brd[col][row]

	visited[[2]int{col, row}] = true
	path = append(path, c)
	nn, ok := node.Next[c]

	if !ok || nn == nil {
		return
	}

	DFs(brd, res, nn, col+1, row, visited, path)
	DFs(brd, res, nn, col-1, row, visited, path)
	DFs(brd, res, nn, col, row+1, visited, path)
	DFs(brd, res, nn, col, row-1, visited, path)
	visited[[2]int{col, row}] = false
	path = path[:len(path)-1]
}
