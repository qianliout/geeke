package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

/*
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	index := findIndex(inorder, postorder[len(postorder)-1])
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func findIndex(order []int, val int) int {
	for i := range order {
		if order[i] == val {
			return i
		}
	}
	return -1
}
