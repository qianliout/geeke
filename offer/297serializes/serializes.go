package main

import (
	"fmt"
	"strconv"
	"strings"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	c := Constructor()
	s := c.serialize(root)
	fmt.Println("res is ", s)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	Fill string
}

func Constructor() Codec {
	return Codec{Fill: "null"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return bfsPreSerialize(root, this.Fill)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	queue := strings.Split(data, ",")
	return dfsIndeserialize(&queue, this.Fill)
}

// dfs前序
func dfsPreserialize(root *TreeNode, fill string) string {
	if root == nil {
		return fill
	}
	left := dfsPreserialize(root.Left, fill)
	right := dfsPreserialize(root.Right, fill)
	return strconv.Itoa(root.Val) + "," + left + "," + right
}

// dfs中序
func dfsInserialize(root *TreeNode, fill string) string {
	if root == nil {
		return fill
	}
	left := dfsPreserialize(root.Left, fill)
	right := dfsPreserialize(root.Right, fill)
	return left + "," + strconv.Itoa(root.Val) + "," + right
}

// dfs后序
func dfsPostserialize(root *TreeNode, fill string) string {
	if root == nil {
		return fill
	}
	left := dfsPreserialize(root.Left, fill)
	right := dfsPreserialize(root.Right, fill)
	return left + "," + right + "," + strconv.Itoa(root.Val)
}

func dfsPredeserialize(queue *[]string, fill string) *TreeNode {

	if len(*queue) == 0 {
		return nil
	}

	i, err := strconv.Atoi((*queue)[0])
	*queue = (*queue)[1:]
	if err != nil {
		return nil
	}

	left := dfsPredeserialize(queue, fill)
	root := &TreeNode{Val: i}
	root.Left = left
	root.Right = dfsPredeserialize(queue, fill)
	return root
}

func dfsIndeserialize(queue *[]string, fill string) *TreeNode {
	left := dfsPredeserialize(queue, fill)

	// 这里一定要放在这里进行判断,不然就会出错
	if len(*queue) == 0 {
		return nil
	}

	i, err := strconv.Atoi((*queue)[0])
	*queue = (*queue)[1:]
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: i}
	root.Left = left

	root.Right = dfsPredeserialize(queue, fill)
	return root
}

func dfsPostdeserialize(queue *[]string, fill string) *TreeNode {
	left := dfsPredeserialize(queue, fill)
	right := dfsPredeserialize(queue, fill)
	// 这里一定要放在这里进行判断,不然就会出错
	if len(*queue) == 0 {
		return nil
	}

	i, err := strconv.Atoi((*queue)[0])
	*queue = (*queue)[1:]
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: i}
	root.Left = left
	root.Right = right
	return root
}

// bfs 层序遍历
func bfsSequenceSerialize(root *TreeNode, fill string) string {
	if root == nil {
		return fill
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	stringList := make([]string, 0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			stringList = append(stringList, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			stringList = append(stringList, fill)
		}
	}
	return strings.Join(stringList, ",")
}

func bfsSequenceDeserialize(ss []string, fill string) *TreeNode {
	if len(ss) == 0 {
		return nil
	}
	idx := 0
	i, err := strconv.Atoi(ss[idx])
	if err != nil { // 这一步一定要判断,不然当输入空结点时会报错
		return nil
	}
	root := &TreeNode{Val: i}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		idx++
		node := queue[0]
		queue = queue[1:]

		if ss[idx] != fill {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				left := &TreeNode{Val: j}
				node.Left = left
				queue = append(queue, left)
			}
		}
		idx++

		if ss[idx] != fill {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				right := &TreeNode{Val: j}
				node.Right = right
				queue = append(queue, right)
			}
		}
	}
	return root
}

// bfs 前序序列化
func bfsPreSerialize(root *TreeNode, fill string) string {
	strSlice := make([]string, 0)
	if root == nil {
		return ""
	}
	stark := []*TreeNode{root}
	for len(stark) > 0 {
		node := stark[len(stark)-1]
		stark = stark[:len(stark)-1]
		if node != nil {
			strSlice = append(strSlice, strconv.Itoa(node.Val))
			stark = append(stark, node.Right)
			stark = append(stark, node.Left)
		} else {
			strSlice = append(strSlice, fill)
		}
	}
	return strings.Join(strSlice, ",")
}

func bfsPreDserialize(ss []string, fill string) *TreeNode {
	if len(ss) == 0 {
		return nil
	}
	idx := 0
	n, err := strconv.Atoi(ss[idx])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: n}
	stark := []*TreeNode{root}
	// 1,2,null,null,3,4,null,null,5,null,null
	for len(stark) > 0 {
		idx++
		node := stark[len(stark)]
		stark = stark[:len(stark)-1]

		if ss[idx] != fill {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				left := &TreeNode{Val: j}
				node.Left = left
				stark = append(stark, left)
			}
		}
		idx++

		if ss[idx] != fill {
			if j, e := strconv.Atoi(ss[idx]); e == nil {
				right := &TreeNode{Val: j}
				node.Right = right
				queue = append(queue, right)
			}
		}

	}
	return root
}
