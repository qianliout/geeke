package main

import (
	. "qianliout/leetcode/common/treenode"
)

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	left := count(root.Left)
	if left == k-1 {
		return root.Val
	} else if left > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-1-left)
	}
}

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root.Left) + count(root.Right) + 1
}
