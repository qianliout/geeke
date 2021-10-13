package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {

}

func isValidBST(root *treenode.TreeNode) bool {

}

func help(min, max int, root *treenode.TreeNode) {

}

// 这种写法有问题，这里的递归是要去改pre,这个值，那么这个值就应该是一个包变量，这样是很不好的，所以不推荐这种写法
func IsValideByInorder(pre, root *treenode.TreeNode) bool {
	if root == nil {
		return true
	}
	// 先遍历左
	if !IsValideByInorder(pre, root.Left) {
		return false
	}
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	if !IsValideByInorder(pre, root.Right) {
		return false
	}
	return true
}

func dfs(root, min, max *treenode.TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && min.Val > root.Val {
		return false
	}
	if max != nil && max.Val < root.Val {
		return false
	}

	return dfs(root.Left, min, root) && dfs(root.Right, root, max)
}

// 这种写法也是可以的
func dfs1(root, min, max *treenode.TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs1(root.Left, min, root) {
		return false
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return dfs1(root.Right, root, max)
}
