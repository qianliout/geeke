package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {

}

/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
*/
func maxDepth(root *TreeNode) int {
	var max int
	dfs(root, &max)
	return max
}
func dfs(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, max)
	if left+1 > *max {
		*max = left + 1
	}

	right := dfs(root.Right, max)
	if right+1 > *max {
		*max = right + 1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}
