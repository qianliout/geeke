package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return dfs(root, root.Val)
}

func dfs(root *TreeNode, min int) int {
	if root == nil {
		return -1
	}
	if root.Val > min {
		return -1
	}
	if root.Val > min {
		return root.Val
	}
	left := dfs(root.Left, min)
	right := dfs(root.Right, min)
	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}
	if right < left {
		return right
	}
	return left
}
