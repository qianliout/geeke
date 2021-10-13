package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {

}

func kthSmallest(root *treenode.TreeNode, k int) int {
	if k == 0 && root != nil {
		return root.Val
	}

}
