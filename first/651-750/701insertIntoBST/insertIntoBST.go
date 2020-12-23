package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val}
				break
			} else {
				p = p.Left
			}
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val}
				break
			} else {
				p = p.Right
			}
		}
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST2(root.Right, val)
	} else {
		root.Left = insertIntoBST2(root.Left, val)
	}
	return root
}
