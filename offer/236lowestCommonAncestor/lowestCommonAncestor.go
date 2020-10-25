package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	var res *TreeNode
	if left != nil && right == nil {
		res = left
	} else if left == nil && right != nil {
		res = right
	} else {
		res = root
	}
	return res
}
