package main

import (
	treenode2 "qianliout/leetcode/common/treenode"
)

func main() {

}

func insertIntoBST(root *treenode2.TreeNode, val int) *treenode2.TreeNode {
	if root == nil {
		return &treenode2.TreeNode{Val: val}
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &treenode2.TreeNode{Val: val}
				break
			} else {
				p = p.Left
			}
		} else {
			if p.Right == nil {
				p.Right = &treenode2.TreeNode{Val: val}
				break
			} else {
				p = p.Right
			}
		}
	}
	return root
}

func insertIntoBST2(root *treenode2.TreeNode, val int) *treenode2.TreeNode {
	if root == nil {
		return &treenode2.TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST2(root.Right, val)
	} else {
		root.Left = insertIntoBST2(root.Left, val)
	}
	return root
}
