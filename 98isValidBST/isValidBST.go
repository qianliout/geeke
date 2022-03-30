package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
*/
func isValidBST(root *TreeNode) bool {
	return dsf(root, nil, nil)
}

func dsf(root *TreeNode, max, min *TreeNode) bool {
	if root == nil {
		return true
	}
	// 一定是>=和<=
	if max != nil && root.Val >= max.Val {
		return false
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	// 先遍历左边还是先遍历右边都是一样的
	return dsf(root.Right, max, root) && dsf(root.Left, root, min)
}
