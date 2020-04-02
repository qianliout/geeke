package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}
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

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	used := make(map[*TreeNode]bool)
	path := make([]int, 0)
	help(root, used, path, &res, sum)
	return res
}

func help(root *TreeNode, used map[*TreeNode]bool, path []int, res *[][]int, n int) {
	// 注意，树中可能有负数，所以不能对n进行小于0的判断
	if root == nil {
		return
	}

	if !used[root] {
		used[root] = true
		path = append(path, root.Val)
		//fmt.Println("path is ", path)
		if n-root.Val == 0 && (root.Left == nil && root.Right == nil) {
			*res = append(*res, append([]int{}, path...))
			return
		}
		help(root.Left, used, path, res, n-root.Val)
		help(root.Right, used, path, res, n-root.Val)
		used[root] = false
		path = path[:len(path)-1]
	}
}
