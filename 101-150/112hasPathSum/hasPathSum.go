package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	res := hasPathSum(root, 1)
	fmt.Println(res)

}

/*
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1
*/
// 注意，这里的数字可能是负数
func hasPathSum(root *TreeNode, sum int) bool {
	res := make([]bool, 0)
	used := make(map[*TreeNode]bool)
	dfs(root, used, sum, &res)
	if len(res) > 0 {
		return true
	}

	return false
}

func dfs(root *TreeNode, used map[*TreeNode]bool, path int, res *[]bool) {
	if len(*res) > 0 {
		return
	}

	if root != nil && path-root.Val == 0 && root.Left == nil && root.Right == nil {
		*res = append(*res, true)
	}
	if !used[root] && root != nil {
		used[root] = true
		path = path - root.Val
		dfs(root.Left, used, path, res)
		dfs(root.Right, used, path, res)
		used[root] = false
		path = path + root.Val
	}
}
