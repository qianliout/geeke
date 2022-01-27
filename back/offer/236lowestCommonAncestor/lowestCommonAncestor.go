package main

import (
	"qianliout/leetcode/back/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	var res *treenode.TreeNode
	if left != nil && right != nil {
		res = root
	}
	if left != nil && right == nil {
		res = left
	}
	if left == nil && right != nil {
		res = right
	}
	return res
}

func lowestCommonAncestor2(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)

	// 左边找到了，右边没有找到，返回左边
	if left != nil && right == nil {
		return left
	}
	// 左边没有找到，右边找到了返回右边
	if left == nil && right != nil {
		return right
	}
	// 两边都找到了，就返回root,说明，
	if left != nil && right != nil {
		return root
	}
	// 都没有找到，就返回空
	return nil
}

// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
func lowestCommonAncestorBST(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if p.Val > q.Val {
		return lowestCommonAncestorBST(root, q, p)
	}
	if root == nil || root == p || root == q {
		return root
	}
	if root.Val < p.Val {
		return lowestCommonAncestorBST(root.Right, p, q)
	}
	if root.Val > q.Val {
		return lowestCommonAncestorBST(root.Left, p, q)
	}
	return root
}
