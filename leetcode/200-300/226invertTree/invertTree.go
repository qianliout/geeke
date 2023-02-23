package main

import (
	"fmt"
	"strings"

	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {
	severity := strings.Split("", ",")
	fmt.Println(len(severity), severity, severity[0])
}

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
