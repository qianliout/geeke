package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	if k == 0 && root != nil {
		return root.Val
	}

}
