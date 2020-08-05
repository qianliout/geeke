package main

import (
	. "outback/leetcode/common"
	. "outback/leetcode/common/treenode"
)

func main() {

}

func maxPathSum(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	*res = Max(*res, root.Val+left, root.Val+right, root.Val, root.Val+left+right)
	return *res
}
