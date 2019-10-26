package tree

import (
	"outback/leetcode/common/treenode"
)

/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：

    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

func IsValidBST(root *treenode.TreeNode) bool {
	return IsValideByRecusion(root, nil, nil)
}

// BST有个特点是，进行一次中序遍历后，是一个升序序列，所以，可以使用。

// 递归的方法
func IsValideByRecusion(root, min, max *treenode.TreeNode) bool {
	// min 记录的是当前最小的那个结点，max记录的当前最大的那个结点，这是关键，不是左结点和右结点
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return IsValideByRecusion(root.Left, min, root) && IsValideByRecusion(root.Right, root, max)
}
