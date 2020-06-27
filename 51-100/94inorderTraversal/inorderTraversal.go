package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	res := inorderTraversal(&root)
	fmt.Println("res is ", res)
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	Iter(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root != nil {
		inorder(root.Left, res)
		*res = append(*res, root.Val)
		inorder(root.Right, res)
	}
}

func Iter(root *TreeNode, res *[]int) {
	stack := make([]*TreeNode, 0)
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
