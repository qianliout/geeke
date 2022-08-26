package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {
}

/*
给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
每条从根节点到叶节点的路径都代表一个数字：
例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
计算从根节点到叶节点生成的 所有数字之和 。
叶节点 是指没有子节点的节点。
*/

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, pre int) int {
	if root == nil {
		return 0
	}
	tmp := pre*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return tmp
	}
	return dfs(root.Left, tmp) + dfs(root.Right, tmp)
}
