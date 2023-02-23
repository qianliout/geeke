package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	ans := getMinimumDifference(root)
	fmt.Println("ans is ", ans)
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := math.MaxInt
	dfs(root, nil, &res)
	return res
}

func dfs(root, pre *TreeNode, res *int) {
	if root == nil {
		return
	}
	dfs(root.Left, root, res)
	if pre != nil && pre.Val-root.Val < *res {
		*res = pre.Val - root.Val
	}
	dfs(root.Right, root, res)
}

func bfs(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	pre, res := math.MinInt, math.MaxInt
	for len(queue) > 0 || root != nil {
		for root != nil {
			queue = append(queue, root)
			root = root.Left
		}
		root = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if pre != math.MinInt && root.Val-pre < res {
			res = root.Val - pre
		}
		pre = root.Val
		root = root.Right
	}
	return res
}
