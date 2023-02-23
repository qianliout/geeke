package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	dfs(root.Right, res)
	*res = append(*res, root.Val)
}

// 后序遍历
func iterate(root *TreeNode) []int {
	res := make([]int, 0)
	stack1, stack2 := make([]*TreeNode, 0), make([]*TreeNode, 0)
	if root != nil {
		stack1 = append(stack1, root)
	}
	for len(stack1) > 0 {
		last := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, last)
		if last.Left != nil {
			stack1 = append(stack1, last.Left)
		}
		if last.Right != nil {
			stack1 = append(stack1, last.Right)
		}
	}
	for i := len(stack2) - 1; i >= 0; i-- {
		res = append(res, stack2[i].Val)
	}
	return res
}
