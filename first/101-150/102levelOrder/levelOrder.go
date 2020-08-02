package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	//root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	res := levelOrder(root)
	fmt.Println("res is ", res)
}

/*
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
*/
func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)
	if root == nil {
		return res
	}
	first := make([]*TreeNode, 0)
	second := make([]*TreeNode, 0)
	first = append(first, root)
	for len(first) > 0 || len(second) > 0 {
		r := make([]int, 0)
		for len(first) > 0 {
			f := first[0]
			first = first[1:len(first)]
			r = append(r, f.Val)
			if f.Left != nil {
				second = append(second, f.Left)
			}
			if f.Right != nil {
				second = append(second, f.Right)
			}

		}
		res = append(res, r)
		first = second
		second = make([]*TreeNode, 0)
	}
	return res
}
