package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。
*/

func flatten(root *TreeNode) {
	help(root)
}

// 展平root，并返回根节点
func help(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	right := help(root.Right)
	left := help(root.Left)
	root.Right = left
	root.Left = nil
	cur := root
	for cur != nil && cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = right
	return root
}
