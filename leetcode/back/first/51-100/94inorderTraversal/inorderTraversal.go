package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := treenode.TreeNode{Val: 1}
	root.Right = &treenode.TreeNode{Val: 2}
	root.Right.Left = &treenode.TreeNode{Val: 3}
	res := inorderTraversal(&root)
	fmt.Println("res is ", res)
}

/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
func inorderTraversal(root *treenode.TreeNode) []int {
	res := make([]int, 0)
	Iter(root, &res)
	return res
}

func inorder(root *treenode.TreeNode, res *[]int) {
	if root != nil {
		inorder(root.Left, res)
		*res = append(*res, root.Val)
		inorder(root.Right, res)
	}
}

func Iter(root *treenode.TreeNode, res *[]int) {
	stack := make([]*treenode.TreeNode, 0)
	curr := root
	for curr != nil || len(stack) > 0 {

		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*res = append(*res, curr.Val)
		curr = curr.Right
	}
}
