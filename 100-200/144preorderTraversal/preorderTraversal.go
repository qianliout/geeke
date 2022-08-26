package main

import (
	"fmt"

	. "qianliout/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 6}
	ans := preorderTraversal(root)
	fmt.Println("ans is ", ans)
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	iteration(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}

func iteration(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*res = append(*res, last.Val)
		if last.Right != nil {
			stack = append(stack, last.Right)
		}
		if last.Left != nil {
			stack = append(stack, last.Left)
		}
	}
}
