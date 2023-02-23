package main

import (
	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
叶子节点 是指没有子节点的节点。
*/
func pathSum(root *TreeNode, targetSum int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs(root, targetSum, path, &res)
	return res
}

func dfs(root *TreeNode, targetSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		*res = append(*res, append([]int{}, path...))
		return
	}
	path = append(path, root.Val)
	dfs(root.Left, targetSum-root.Val, path, res)
	dfs(root.Right, targetSum-root.Val, path, res)
	path = path[:len(path)-1]
}
