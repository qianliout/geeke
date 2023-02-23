package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {

}

/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := root.Val
	dfs(root, &max)
	return max
}

func dfs(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, max)
	right := dfs(root.Right, max)
	*max = Max(*max, root.Val, root.Val+left, root.Val+right, root.Val+left+right)
	return Max(root.Val, root.Val+left, root.Val+right)
}
