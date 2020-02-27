package tree

import (
	"outback/leetcode/common/treenode"
)

/*
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
示例 1:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:
输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
*/

func LowestCommonAncestorBST(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	// 二叉搜索树也是二叉树，所以235的做法仍然也是有用的
	// 因为是二叉搜索树，所以可以加快速度
	if p.Val < root.Val && q.Val < root.Val {
		// 说明p q 都在root的左边
		return LowestCommonAncestorBST(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		// 说明p q 都在root的右边
		return LowestCommonAncestorBST(root.Right, p, q)
	}
	// 如果p q 在root的两边。就返回root
	return root
}

func LowestCommonAncestorBSTByWhile(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
			continue
		}
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
			continue
		}
		return root
	}
	return root
}
