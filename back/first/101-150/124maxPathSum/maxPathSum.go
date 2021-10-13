package main

import (
	"fmt"

	"outback/leetcode/back/common"
	"outback/leetcode/back/common/treenode"
)

func main() {

	// root := &TreeNode{Val: -10}
	// root.Left = &TreeNode{Val: 9}
	// root.Right = &TreeNode{Val: 20}
	// root.Right.Left = &TreeNode{Val: 15}
	// root.Right.Right = &TreeNode{Val: 7}

	root := &treenode.TreeNode{Val: 1}
	root.Left = &treenode.TreeNode{Val: -1}
	root.Right = &treenode.TreeNode{Val: 3}
	res := maxPathSum(root)
	fmt.Println("res is ", res)
}

/*
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
示例 1:
输入: [1,2,3]
       1
      / \
     2   3
输出: 6
示例 2:
输入: [-10,9,20,null,null,15,7]
   -10
   / \
  9  20
    /  \
   15   7
输出: 42
*/
// var res int

func maxPathSum(root *treenode.TreeNode) int {
	// 一定要先赋值，可能出现的一种情况是，只有root一个结点，并且值为负，
	res := root.Val
	Dfs(root, &res)
	return res
}

// Dfs 的返回值表示走到root这个节点并选中root这个节点的最大值
func Dfs(root *treenode.TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := Dfs(root.Left, res)
	right := Dfs(root.Right, res)
	*res = common.Max(*res, root.Val+left+right, root.Val+left, root.Val+right, root.Val)
	return common.Max(root.Val+left, root.Val+right, root.Val)
}
