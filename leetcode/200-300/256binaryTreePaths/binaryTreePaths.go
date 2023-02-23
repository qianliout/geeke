package main

import (
	"fmt"
	"strings"

	. "qianliout/leetcode/leetcode/common/treenode"
)

func main() {

}

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径
*/
func binaryTreePaths(root *TreeNode) []string {
	path := make([]int, 0)
	res := make([]string, 0)
	dsf(root, path, &res)
	return res
}

func dsf(root *TreeNode, path []int, res *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		*res = append(*res, convert(path))
		return
	}
	path = append(path, root.Val)
	if root.Left != nil {
		dsf(root.Left, path, res)
	}
	if root.Right != nil {
		dsf(root.Right, path, res)
	}
}

func convert(res []int) string {
	ans := make([]string, 0)
	for i := range res {
		ans = append(ans, fmt.Sprintf("%d", res[i]))
	}
	return strings.Join(ans, "->")
}

func binaryTreePaths2(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, fmt.Sprintf("%d", root.Val))
		return res
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	for i := range left {
		res = append(res, fmt.Sprintf("%d->%s", root.Val, left[i]))
	}
	for i := range right {
		res = append(res, fmt.Sprintf("%d->%s", root.Val, right[i]))
	}

	return res
}
