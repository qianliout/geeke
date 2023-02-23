package main

func main() {

}

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

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{Root: NewNode()}
}

func (w *WordDictionary) AddWord(word string) {
	w.add(w.Root, []byte(word), 0)
}
func (w *WordDictionary) add(root *Node, word []byte, index int) {
	if index >= len(word) {
		root.End = true
		return
	}
	c := word[index]
	if root.Next[c] == nil {
		root.Next[c] = NewNode()
	}
	w.add(root.Next[c], word, index+1)
}

func (w *WordDictionary) find(root *Node, word []byte, index int) bool {
	if index >= len(word) {
		return root.End
	}

	if root.Next == nil {
		return false
	}

	c := word[index]
	if c == '.' {
		for _, v := range root.Next {
			if v != nil && w.find(v, word, index+1) {
				return true
			}
		}
		return false
	}

	if root.Next[c] == nil {
		return false
	}
	return w.find(root.Next[c], word, index+1)

}

func (w *WordDictionary) Search(word string) bool {
	return w.find(w.Root, []byte(word), 0)
}
