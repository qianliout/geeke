package main

import (
	"fmt"
	"strconv"
	"strings"

	treenode2 "qianliout/leetcode/leetcode/common/treenode"
)

func main() {
	root := &treenode2.TreeNode{Val: 1}
	root.Left = &treenode2.TreeNode{Val: 2}
	root.Right = &treenode2.TreeNode{Val: 3}
	root.Right.Left = &treenode2.TreeNode{Val: 4}
	root.Right.Right = &treenode2.TreeNode{Val: 5}
	c := Constructor()
	s := c.serialize(root)

	node := c.deserialize(s)
	fmt.Println("serialize", s)
	treenode2.PreOrderTraversal(node)
}

/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，
你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
示例:
你可以将以下二叉树：
    1
   / \
  2   3
     / \
    4   5
序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，
你也可以采用其他的方法解决这个问题。
说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
*/

type Codec struct {
	Zero string
}

func Constructor() Codec {
	return Codec{Zero: "null"}
}

// 这道题,最后的几个Null可以不用去除,这样会好的多,这个就是一个完全二叉树,反序列化时就会好的很多,而且,也不用加前后的[ ]
// Serializes a tree to a single string.
func (this *Codec) serialize(root *treenode2.TreeNode) string {
	// return rserialize(root, "")
	return rserialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *treenode2.TreeNode {
	// 先生成数组
	queue := strings.Split(data, ",")
	// bfs
	// root := rdeserialize(&queue)
	root := bfsDeserialize(&queue)
	return root
}

// bfs
func bfsSerialize(root *treenode2.TreeNode) string {
	if root == nil {
		return "null"
	}
	s := make([]string, 0)
	queue := make([]*treenode2.TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if first == nil {
			s = append(s, "null")
		} else {
			s = append(s, strconv.Itoa(first.Val))
			queue = append(queue, first.Left)
			queue = append(queue, first.Right)
		}
	}
	return strings.Join(s, ",")
}

func bfsDeserialize(ss *[]string) *treenode2.TreeNode {
	s := *ss
	idx := 0
	i, err := strconv.Atoi(s[idx])

	if err != nil {
		return nil
	}
	root := &treenode2.TreeNode{Val: i}

	queue := []*treenode2.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		idx++

		if s[idx] != "null" {
			if j, e := strconv.Atoi(s[idx]); e == nil {
				left := &treenode2.TreeNode{Val: j}
				node.Left = left
				queue = append(queue, left)
			}
		}
		idx++
		if s[idx] != "null" {
			if j, e := strconv.Atoi(s[idx]); e == nil {
				right := &treenode2.TreeNode{Val: j}
				node.Right = right
				queue = append(queue, right)
			}
		}
	}
	return root
}

// 使用递归的方式
func rdeserialize(queue *[]string) *treenode2.TreeNode {
	if len(*queue) == 0 {
		return nil
	}

	i, err := strconv.Atoi((*queue)[0])
	*queue = (*queue)[1:]
	if err != nil {
		return nil
	}
	root := &treenode2.TreeNode{Val: i}
	root.Left = rdeserialize(queue)
	root.Right = rdeserialize(queue)
	return root
}

func rserialize(root *treenode2.TreeNode, s string) string {
	if root == nil {
		s = s + "null,"
	} else {
		s = s + strconv.Itoa(root.Val) + ","
		s = rserialize(root.Left, s)
		s = rserialize(root.Right, s)
	}
	return s
}
