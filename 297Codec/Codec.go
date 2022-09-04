package main

import (
	"fmt"
	"strconv"
	"strings"

	. "qianliout/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	codec := Constructor()
	ser := codec.serialize(root)
	node := codec.deserialize(ser)

	PreOrderTraversal(node)

	fmt.Println("   ser is ", ser)

}

/*
二叉树的序列化与反序列化

*/

// 层序
type CodecLevel struct {
	Separate string
	Null     string
}

func ConstructorLevel() CodecLevel {
	return CodecLevel{Separate: ",", Null: "null"}
}

// Serializes a tree to a single string.
func (this *CodecLevel) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ans := make([]string, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root)
	for len(que) > 0 {
		tmp := make([]*TreeNode, 0)
		for i := range que {
			if que[i] != nil {
				ans = append(ans, fmt.Sprintf("%d", que[i].Val))
				tmp = append(tmp, que[i].Left)
				tmp = append(tmp, que[i].Right)
			} else {
				ans = append(ans, this.Null)
			}
		}
		que = tmp
	}
	return strings.Join(ans, this.Separate)
}

// Deserializes your encoded data to tree.
func (this *CodecLevel) deserialize(data string) *TreeNode {
	ss := strings.Split(data, this.Separate)
	if len(ss) <= 1 || ss[0] == this.Null || ss[0] == "" {
		return nil
	}
	que := make([]*TreeNode, 0)
	n, _ := strconv.Atoi(ss[0])
	root := &TreeNode{Val: n}
	que = append(que, root)
	j, i := 0, 1
	for len(que) > 0 {
		if i >= len(ss) {
			break
		}
		first := que[j]
		j++
		// 左
		if ss[i] != this.Null {
			n, _ = strconv.Atoi(ss[i])
			left := &TreeNode{Val: n}
			que = append(que, left)
			first.Left = left
		}
		i++
		// 右
		if ss[i] != this.Null {
			n, _ = strconv.Atoi(ss[i])
			right := &TreeNode{Val: n}
			que = append(que, right)
			first.Right = right
		}
		i++
	}
	return root
}
