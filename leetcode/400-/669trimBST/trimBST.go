package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

func trimBST1(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST1(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST1(root.Left, low, high)
	}
	root.Left = trimBST1(root.Left, low, high)
	root.Right = trimBST1(root.Right, low, high)
	return root
}

func trimBST(root *TreeNode, low, high int) *TreeNode {
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else if root.Val > high {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	// 左边
	for node := root; node.Left != nil; {
		if node.Left.Val < low {
			node.Left = node.Left.Right
		} else {
			node = node.Left
		}
	}
	for node := root; node.Right != nil; {
		if node.Right.Val > high {
			node.Right = node.Right.Left
		} else {
			node = node.Right
		}
	}
	return root
}
