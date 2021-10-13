package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {

}

func insertIntoBST(root *treenode.TreeNode, val int) *treenode.TreeNode {
	if root == nil {
		return &treenode.TreeNode{Val: val}
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &treenode.TreeNode{Val: val}
				break
			} else {
				p = p.Left
			}
		} else {
			if p.Right == nil {
				p.Right = &treenode.TreeNode{Val: val}
				break
			} else {
				p = p.Right
			}
		}
	}
	return root
}

func insertIntoBST2(root *treenode.TreeNode, val int) *treenode.TreeNode {
	if root == nil {
		return &treenode.TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST2(root.Right, val)
	} else {
		root.Left = insertIntoBST2(root.Left, val)
	}
	return root
}
