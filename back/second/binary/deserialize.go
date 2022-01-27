package main

import (
	"fmt"
	"strconv"
	"strings"

	"qianliout/leetcode/back/common/treenode"
)

func main() {
	root := &treenode.TreeNode{Val: 1}
	root.Left = &treenode.TreeNode{Val: 2}
	root.Right = &treenode.TreeNode{Val: 3}
	root.Right.Left = &treenode.TreeNode{Val: 4}
	root.Right.Right = &treenode.TreeNode{Val: 5}
	c := Constructor()
	s := c.serialize(root)
	fmt.Println("res is ", s)
	node := c.deserialize(s)
	treenode.PreOrderTraversal(node)
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
	fill string
}

func Constructor() Codec {
	return Codec{fill: "null"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *treenode.TreeNode) string {
	return BFSs(root, this.fill)
}

func DFSs(root *treenode.TreeNode, fill string) string {
	if root == nil {
		return fill
	}
	return strconv.Itoa(root.Val) + "," + DFSs(root.Left, fill) + "," + DFSs(root.Right, fill)
}

// 其实就是层序层序遍历
func BFSs(root *treenode.TreeNode, fill string) string {
	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	ans := ""
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		if first == nil {
			ans = ans + fill + ","
		} else {
			ans = ans + strconv.Itoa(first.Val) + ","
		}
		if first != nil {
			queue = append(queue, first.Left)
			queue = append(queue, first.Right)
		}
	}
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *treenode.TreeNode {
	list := strings.Split(data, ",")
	return BFSd(list)
}

func DFSd(data *[]string) *treenode.TreeNode {
	if len(*data) == 0 {
		return nil
	}
	v, err := strconv.Atoi((*data)[0])
	// 这上不一定要先做再返回，不然就会出错，怎么理解呢，就也是你这个值取出来，
	// 发现是一个nil结点，但是不管怎么样，这个元素都应该从list去掉
	*data = (*data)[1:]
	if err != nil {
		return nil
	}
	root := &treenode.TreeNode{Val: v}
	root.Left = DFSd(data)
	root.Right = DFSd(data)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func BFSd(data []string) *treenode.TreeNode {
	s := data
	idx := 0
	i, err := strconv.Atoi(s[idx])

	if err != nil {
		return nil
	}
	root := &treenode.TreeNode{Val: i}

	queue := []*treenode.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		idx++

		if s[idx] != "null" {
			if j, e := strconv.Atoi(s[idx]); e == nil {
				left := &treenode.TreeNode{Val: j}
				node.Left = left
				queue = append(queue, left)
			}
		}
		idx++
		if s[idx] != "null" {
			if j, e := strconv.Atoi(s[idx]); e == nil {
				right := &treenode.TreeNode{Val: j}
				node.Right = right
				queue = append(queue, right)
			}
		}
	}
	return root

}
