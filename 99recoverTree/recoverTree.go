package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	recoverTree(root)
	InOrderTraversal(root)
}

/*
给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树
*/

func recoverTree(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else { // 这里要在下一个循环才退出，因为下一个循环可能改变y的值,但是不改变x的值
				break
			}
		}
		pred = root
		root = root.Right
	}
	if x != nil && y != nil {
		x.Val, y.Val = y.Val, x.Val
	}
}
