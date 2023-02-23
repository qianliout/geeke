package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
*/
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root != nil && root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	if root != nil {
		res = append(res, root.Val)
	}
	if root != nil && root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

func iterate(root *TreeNode) []int {
	res := make([]int, 0)
	stark := make([]*TreeNode, 0)

	for len(stark) > 0 || root != nil {
		if root != nil {
			stark = append(stark, root)
			root = root.Left
		} else {
			last := stark[len(stark)-1]
			stark = stark[:len(stark)-1]
			res = append(res, last.Val)
			root = last.Right
		}
	}
	return res
}
