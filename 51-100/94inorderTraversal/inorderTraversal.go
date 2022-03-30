package main

import (
	. "qianliout/leetcode/common/treenode"
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
