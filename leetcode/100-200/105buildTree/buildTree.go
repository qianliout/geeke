package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	InOrderTraversal(root)
	PreOrderTraversal(root)
}

/*
从前序与中序遍历序列构造二叉树
这一题目可以假设没有重复值
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	index := findIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
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
