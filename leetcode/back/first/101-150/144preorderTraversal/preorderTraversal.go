package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := &treenode.TreeNode{Val: 1}
	root.Left = &treenode.TreeNode{Val: 4}
	root.Left.Left = &treenode.TreeNode{Val: 2}
	root.Left.Right = &treenode.TreeNode{Val: 5}
	root.Right = &treenode.TreeNode{Val: 3}
	res := preorderTraversal(root)
	fmt.Println("res is ", res)
}

/*
给定一个二叉树，返回它的 前序 遍历。
 示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
*/
func preorderTraversal(root *treenode.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	//pre(root, &res)
	Iter(root, &res)
	return res
}

// 递归解法
func pre(root *treenode.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	if root.Left != nil {
		pre(root.Left, res)
	}
	if root.Right != nil {
		pre(root.Right, res)
	}
}

// 迭代解法,这里一技巧
func Iter(root *treenode.TreeNode, res *[]int) {
	stack := make([]*treenode.TreeNode, 0)
	if root != nil {
		stack = append(stack, root)
	}
	for len(stack) > 0 {
		first := stack[len(stack)-1] // 从后从前取,所以先加right,再加left
		stack = stack[:len(stack)-1]
		*res = append(*res, first.Val)

		if first.Right != nil {
			stack = append(stack, first.Right)
		}
		if first.Left != nil {
			stack = append(stack, first.Left)
		}
	}
}
