package main

import (
	"strconv"

	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
两棵树重复是指它们具有相同的结构以及相同的结点值。
*/

// 使用序列化的方式
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)
	exit := make(map[string]int)
	visit := make(map[*TreeNode]bool)

	dfs(root, &res, exit, visit)

	return res
}

func dfs(root *TreeNode, res *[]*TreeNode, exit map[string]int, visit map[*TreeNode]bool) {
	if root == nil {
		return
	}
	df := serialization(root)
	visit[root] = true
	if exit[df] == 1 {
		*res = append(*res, root)
	}
	exit[df] += 1
	dfs(root.Left, res, exit, visit)
	dfs(root.Right, res, exit, visit)
	visit[root] = false
}

func serialization(node *TreeNode) string {
	if node == nil {
		return "null"
	}
	return strconv.Itoa(node.Val) + "," + serialization(node.Left) + "," + serialization(node.Right)
}
