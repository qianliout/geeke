package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {

}

func isBalanced(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	left := findDeep(root.Left)
	right := findDeep(root.Right)
	if AbsSub(left, right) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func findDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	return Max(findDeep(root.Left), findDeep(root.Right)) + 1
}
