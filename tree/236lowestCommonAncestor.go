package tree

import "outback/leetcode/common/treenode"

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
*/

func LowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if root == p || root == q || root == nil {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	var result *treenode.TreeNode
	// 如果左边找不到，右边找到了，返回右边
	if left == nil && right != nil {
		result = right
	}
	// 如果右边找不到，左边找到了，返回左边
	if left != nil && right == nil {
		result = left
	}
	// 如果两边都找到了，那就返货root
	if left != nil && right != nil {
		result = root
	}
	return result
}
