package main

//
//import (
//	"fmt"
//)
//
//func main() {
//	obj := Constructor()
//	obj.AddWord("bad")
//	obj.AddWord("dad")
//	obj.AddWord("mad")
//	//fmt.Println(int(byte('b')))
//
//	fmt.Println(obj.Query(".ad"))
//}
//
///*
//设计一个支持以下两种操作的数据结构：
//void addWord(word)
//bool search(word)
//search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。
//示例:
//addWord("bad")
//addWord("dad")
//addWord("mad")
//search("pad") -> false
//search("bad") -> true
//search(".ad") -> true
//search("b..") -> true
//链接：https://leetcode-cn.com/problems/add-and-search-word-data-structure-design
//*/
//type Node struct {
//	isWord bool
//	Next   map[string]*Node
//}
//
//type WordDictionary struct {
//	root *Node
//}
//
//func NewNode() *Node {
//	return &Node{Next: make(map[string]*Node)}
//}
//
///** Initialize your data structure here. */
//func Constructor() WordDictionary {
//	return WordDictionary{root: NewNode()}
//}
//
///** Adds a word into the data structure. */
//func (this *WordDictionary) AddWord(word string) {
//	add(this.root, word, 0)
//}
//
//func add(node *Node, word string, index int) {
//	if len(word) == index {
//		node.isWord = true
//		return
//	}
//
//	var c string
//	for {
//		c = string([]rune(word)[index])
//		if c == "." {
//			index++
//		} else {
//			break
//		}
//	}
//	if node.Next[c] == nil {
//		node.Next[c] = NewNode()
//	}
//	add(node.Next[c], word, index+1)
//}
//
///*
//Returns if the word is in the data structure. A word could contain the dot character
//'.' to represent any one letter.
//*/
//func (this *WordDictionary) Query(word string) bool {
//	return find(this.root, word, 0)
//}
//
//func find(node *Node, word string, index int) bool {
//	if len(word) == index {
//		return node.isWord
//	}
//	c := string([]rune(word)[index])
//	if node.Next == nil {
//		return false
//	}
//	if c != "." {
//		if node.Next[c] == nil {
//			return false
//		}
//		return find(node.Next[c], word, index+1)
//	} else {
//		// c=="."的情况下
//		for _, value := range node.Next {
//			if value != nil && find(value, word, index+1) {
//				return true
//			}
//		}
//		return false
//	}
//}
//
//func Or(ss []bool) bool {
//	for _, s := range ss {
//		if s {
//			return true
//		}
//	}
//	return false
//}
//
///**
// * Your WordDictionary object will be instantiated and called as such:
// * obj := Constructor();
// * obj.AddWord(word);
// * param_2 := obj.Query(word);
// */

/*
设计一个支持以下两种操作的数据结构：

void addWord(word)
bool search(word)
search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。

示例:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
链接：https://leetcode-cn.com/problems/add-and-search-word-data-structure-design
*/

type WordDictionary struct {
	Root *Node
}
type Node struct {
	IsWord bool
	Next   map[byte]*Node
}

func newNode() *Node {
	return &Node{Next: make(map[byte]*Node)}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{Root: newNode()}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	add(this.Root, []byte(word), 0)
}

func add(node *Node, word []byte, index int) {

	if index == len(word) {
		node.IsWord = true
		return
	}
	c := word[index]
	if node.Next[c] == nil {
		node.Next[c] = newNode()
	}
	add(node.Next[c], word, index+1)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return find(this.Root, []byte(word), 0)
}

func find(node *Node, word []byte, index int) bool {
	if index == len(word) {
		return node.IsWord
	}
	c := word[index]
	if node.Next == nil {
		return false
	}

	// 通佩服
	if c == '.' {
		for _, value := range node.Next {
			if value != nil && find(value, word, index+1) {
				return true
			}
		}
		return false
	}
	// 这里是易错点,这里不能在通佩服前面判断
	if node.Next[c] == nil {
		return false
	}
	return find(node.Next[c], word, index+1)
}
