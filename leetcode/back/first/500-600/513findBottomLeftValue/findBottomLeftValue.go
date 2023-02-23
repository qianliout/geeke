package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	res := findBottomLeftValue(root)
	fmt.Println("res is ", res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
给定一个二叉树，在树的最后一行找到最左边的值。
示例 1:
输入:

	  2
	 / \
	1   3

输出:
1
示例 2:
输入:

	    1
	   / \
	  2   3
	 /   / \
	4   5   6
	   /
	  7

输出:
7
注意: 您可以假设树（即给定的根节点）不为 NULL。
*/
func findBottomLeftValue(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root.Left != nil && findDepth(root.Left) >= findDepth(root.Right) {
		return dfs(root.Left)
	} else if root.Right != nil && findDepth(root.Left) < findDepth(root.Right) {
		return dfs(root.Right)
	}
	return root.Val
}

// 这里有大量的重复计算，可以用map做优化
func findDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := findDepth(root.Left)
	right := findDepth(root.Right)
	if left >= right {
		return left + 1
	}
	return right + 1
}
