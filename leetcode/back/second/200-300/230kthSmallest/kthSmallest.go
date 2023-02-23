package main

import (
	treenode2 "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

func kthSmallest(root *treenode2.TreeNode, k int) int {
	if k == 0 && root != nil {
		return root.Val
	}

}
