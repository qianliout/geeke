package main

import (
	"container/list"
	. "outback/leetcode/common/treenode"
)

func main() {

}

/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：
[
  [15,7],
  [9,20],
  [3]
]
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
*/
// 想法是和106一样遍历，只不过在结果时反向一次
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	res = order(root)
	result := make([][]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, res[i])
	}
	return result
}

func order(root *TreeNode) [][]int {
	res := make([][]int, 0)
	first := make([]*TreeNode, 0)
	first = append(first, root)
	second := make([]*TreeNode, 0)
	for len(first) > 0 || len(second) > 0 {
		th := make([]int, 0)
		for len(first) > 0 {
			node := first[0]
			first = first[1:]
			th = append(th, node.Val)
			if node.Left != nil {
				second = append(second, node.Left)
			}
			if node.Right != nil {
				second = append(second, node.Right)
			}
		}
		res = append(res, th)
		first = second
		second = make([]*TreeNode, 0)
	}
	return res
}

// 使用go提供的list ,但是后面要输出[][]int的格式时，还是要有一次遍历，时间复杂度和空间复杂度没有变化
func order2(root *TreeNode) *list.List {
	res := list.New()
	first := make([]*TreeNode, 0)
	first = append(first, root)
	second := make([]*TreeNode, 0)
	for len(first) > 0 || len(second) > 0 {
		th := make([]int, 0)
		for len(first) > 0 {
			node := first[0]
			first = first[1:]
			th = append(th, node.Val)
			if node.Left != nil {
				second = append(second, node.Left)
			}
			if node.Right != nil {
				second = append(second, node.Right)
			}
		}
		res.PushFront(th)
		first = second
		second = make([]*TreeNode, 0)
	}
	return res
}
