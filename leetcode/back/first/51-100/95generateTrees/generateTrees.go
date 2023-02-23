package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	res := generateTrees(3)
	fmt.Println("res is ", len(res))
}

/*
 给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
示例:
输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/
func generateTrees(n int) []*treenode.TreeNode {
	if n == 0 {
		return make([]*treenode.TreeNode, 0)
	}
	return helper(1, n)
}

func helper(start, end int) []*treenode.TreeNode {
	// n 表示总数，i 表示第i 做根结点
	res := make([]*treenode.TreeNode, 0)
	// 这个基本条件是怎么判断的呢，为什么要加一个nil呢，要好好理解
	// 结合下面这一名rightNodes := helper(i+1, end),也就是当每次的i=n时，只是左边有数据，
	// 左边没有子树，那么就应该是一个nil,如果不写这个，那么也就不会循环右边的循环了，
	if start > end {
		res = append(res, nil)
		fmt.Println("start end ", start, end, len(res))
		return res
	}
	for i := start; i <= end; i++ {
		leftNodes := helper(start, i-1)
		rightNodes := helper(i+1, end)
		for _, leftNode := range leftNodes {
			for _, rightNode := range rightNodes {
				node := &treenode.TreeNode{Val: i}
				node.Left = leftNode
				node.Right = rightNode
				res = append(res, node)
			}
		}
	}
	return res
}

func dfs(start, end int) []*treenode.TreeNode {
	var res []*treenode.TreeNode
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		left := dfs(start, i-1)
		right := dfs(i+1, end)
		for _, l := range left {
			for _, r := range right {
				node := treenode.TreeNode{Val: i}
				node.Left = l
				node.Right = r
				res = append(res, &node)
			}
		}
	}
	return res
}
