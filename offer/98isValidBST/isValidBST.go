package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

func isValidBST(root *TreeNode) bool {

}

func help(min, max int, root *TreeNode) {

}

func IsValideByInorder(pre, root *TreeNode) bool {
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

func dfs(root, min, max *TreeNode) bool {
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
