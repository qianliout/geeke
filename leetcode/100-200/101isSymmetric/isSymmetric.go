package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return symmetric(left.Right, right.Left) && symmetric(left.Left, right.Right)
}
