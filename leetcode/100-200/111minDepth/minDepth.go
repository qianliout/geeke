package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {

}

/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return Min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}
