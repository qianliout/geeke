package main

import (
	"fmt"
	"strconv"
	"strings"

	. "qianliout/leetcode/common/treenode"

	. "qianliout/leetcode/leetcode/common/treenode"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 前序
type CodecPreOrder struct {
	Separate string
	Null     string
}

func ConstructorPreOrder() CodecPreOrder {
	return CodecPreOrder{Separate: ",", Null: "null"}
}

// Serializes a tree to a single string.
func (this *CodecPreOrder) serialize(root *TreeNode) string {
	if root == nil {
		return this.Null
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	return fmt.Sprintf("%d%s%s%s%s", root.Val, this.Separate, left, this.Separate, right)
}

// Deserializes your encoded data to tree.
func (this *CodecPreOrder) deserialize(data string) *TreeNode {
	ss := strings.Split(data, this.Separate)
	if len(ss) <= 1 || ss[0] == this.Null || ss[0] == "" {
		return nil
	}
	return PreOrderDfs(&ss)
}

func PreOrderDfs(data *[]string) *TreeNode {
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
	root := &TreeNode{Val: v}
	root.Left = PreOrderDfs(data)
	root.Right = PreOrderDfs(data)
	return root
}

/**
 * Your CodecPreOrder object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
