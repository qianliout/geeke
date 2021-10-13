package main

import (
	"outback/leetcode/back/common"
	"outback/leetcode/back/common/treenode"
)

func main() {

}

/*
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。
注意：两个节点之间的路径长度由它们之间的边数表示。
示例 1:
输入:
              5
             / \
            4   5
           / \   \
          1   1   5
输出:
2
示例 2:
输入:
              1
             / \
            4   5
           / \   \
          4   4   5
输出:
2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。
*/
func longestUnivaluePath(root *treenode.TreeNode) int {
	var ans int
	help(root, &ans)
	return ans
}

// 表示以root为根节点，且只是单边的结果
func help(root *treenode.TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := help(root.Left, ans)
	right := help(root.Right, ans)
	var res int
	// 两边相等
	if root.Left != nil && root.Left.Val == root.Val && root.Right != nil && root.Right.Val == root.Val {
		*ans = common.Max(*ans, left+right+2)
	}
	// 左边相等,就先计算左边
	if root.Left != nil && root.Left.Val == root.Val {
		res = left + 1
	}
	// 右边相等，再看右边
	if root.Right != nil && root.Right.Val == root.Val {
		res = common.Max(res, right+1)
	}

	*ans = common.Max(*ans, res)
	return res
}
