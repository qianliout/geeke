package main

import (
	treenode2 "qianliout/leetcode/common/treenode"
	"qianliout/leetcode/common/utils"
)

func main() {

}

/*
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。
示例 1：
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
*/
func maxPathSum(root *treenode2.TreeNode) int {
	// 一定要先赋值。因为可能有负数
	var max int = root.Val

	help(root, &max)
	return max

}

// 返回值表示只走一边且包扩包前节点的最大值，因为只走一边，所以在计算max的值时可以root.Val+left+right
func help(root *treenode2.TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := help(root.Left, max)
	right := help(root.Right, max)
	*max = utils.Max(*max, root.Val, root.Val+left, root.Val+right, root.Val+left+right)
	// 不能这样写的原因是，只走一边，所以上面max才可以左右都相加
	// return Max(root.Val, root.Val+left, root.Val+right,root+right+left)
	return utils.Max(root.Val, root.Val+left, root.Val+right)
}

func help2(root *treenode2.TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := help2(root.Left, max)
	ritght := help2(root.Right, max)
	*max = utils.Max(root.Val, root.Val+left, root.Val+left, root.Val+ritght, root.Val+left+ritght)
	return utils.Max(root.Val, root.Val+left, root.Val+left, root.Val+ritght)
}
