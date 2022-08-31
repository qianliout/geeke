package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}
	if root == nil || root == p || root == q {
		return root
	}
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
