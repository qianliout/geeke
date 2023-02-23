package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := treenode.TreeNode{Val: 5}
	root.Left = &treenode.TreeNode{Val: 4}
	root.Left.Left = &treenode.TreeNode{Val: 11}
	root.Left.Left.Left = &treenode.TreeNode{Val: 7}
	root.Left.Left.Right = &treenode.TreeNode{Val: 2}

	root.Right = &treenode.TreeNode{Val: 8}
	root.Right.Left = &treenode.TreeNode{Val: 13}
	root.Right.Right = &treenode.TreeNode{Val: 4}
	root.Right.Right.Left = &treenode.TreeNode{Val: 5}
	root.Right.Right.Right = &treenode.TreeNode{Val: 1}
	res := pathSum(&root, 22)
	fmt.Println("res is ", res)

}

/*
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和 sum = 22，
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
链接：https://leetcode-cn.com/problems/path-sum-ii
*/

func pathSum(root *treenode.TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	used := make(map[*treenode.TreeNode]bool)
	path := make([]int, 0)
	help(root, used, path, &res, sum)
	return res
}

func help(root *treenode.TreeNode, used map[*treenode.TreeNode]bool, path []int, res *[][]int, n int) {
	// 注意，树中可能有负数，所以不能对n进行小于0的判断
	if root == nil {
		return
	}

	if !used[root] {
		used[root] = true
		path = append(path, root.Val)
		// fmt.Println("path is ", path)
		if n-root.Val == 0 && (root.Left == nil && root.Right == nil) {
			*res = append(*res, append([]int{}, path...))
			return // 这个return可写可以不写，不影响
		}
		help(root.Left, used, path, res, n-root.Val)
		help(root.Right, used, path, res, n-root.Val)
		used[root] = false
		path = path[:len(path)-1]
	}
}
