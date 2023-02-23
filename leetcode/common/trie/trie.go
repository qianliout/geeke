package trie

type Node struct {
	IsWord bool
	Next   map[string]*Node
}

func NewNode() *Node {
	return &Node{Next: make(map[string]*Node)}
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: NewNode()}
}

// 向Trie中添加一个新的单词word
func (t *Trie) Insert(word string) {
	add(t.Root, word, 0)
}

// 向以node为根的Trie中添加word[index...end)，递归算法
func add(node *Node, word string, index int) {
	if index == len(word) {
		node.IsWord = true
		return
	}

	c := string([]rune(word)[index])
	if node.Next[c] == nil {
		node.Next[c] = NewNode()
	}
	add(node.Next[c], word, index+1)
}

// 查询单词word是否在Trie中
func (t *Trie) Search(word string) bool {
	return contains(t.Root, word, 0)
}

// 在以node为根的Trie中查询单词word[index...end)是否存在，递归算法
func contains(node *Node, word string, index int) bool {
	if index == len(word) {
		return node.IsWord
	}

	c := string([]rune(word)[index])
	if node.Next[c] == nil {
		return false
	}
	return contains(node.Next[c], word, index+1)
}

// 查询是否在Trie中有单词以prefix为前缀
func (t *Trie) StartsWith(prefix string) bool {
	return isPrefix(t.Root, prefix, 0)
}

// 查询在以Node为根的Trie中是否有单词以prefix[index...end)为前缀，递归算法
func isPrefix(node *Node, prefix string, index int) bool {
	if index == len(prefix) {
		return true
	}

	c := string([]rune(prefix)[index])
	if node.Next[c] == nil {
		return false
	}

	return isPrefix(node.Next[c], prefix, index+1)
}
