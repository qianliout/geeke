package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	res := zigzagLevelOrder(&root)
	fmt.Println("res is ", res)
}

/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：
[
  [3],
  [20,9],
  [15,7]
]
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	firstStack := make([]*TreeNode, 0)
	secondStack := make([]*TreeNode, 0)
	firstStack = append(firstStack, root)
	for {
		s := make([]int, 0)
		s2 := make([]int, 0)
		for len(firstStack) > 0 {
			node := firstStack[0]
			fmt.Println(node.Val)
			if node.Left != nil {
				secondStack = append(secondStack, node.Left)
			}
			if node.Right != nil {
				secondStack = append(secondStack, node.Right)
			}
			s = append(s, node.Val)
			firstStack = firstStack[1:]
		}
		if len(s) > 0 {
			res = append(res, s)
			s = make([]int, 0)
		}

		for len(secondStack) > 0 {
			node := secondStack[len(secondStack)-1]
			if node.Left != nil {
				firstStack = append(firstStack, node.Left)
			}
			if node.Right != nil {
				firstStack = append(firstStack, node.Right)
			}
			s2 = append(s2, node.Val)
			secondStack = secondStack[:len(secondStack)-1]
		}
		if len(s2) > 0 {
			res = append(res, s2)
			s2 = make([]int, 0)
		}

		if len(firstStack) == 0 && len(secondStack) == 0 {
			break
		}
	}
	return res
}
