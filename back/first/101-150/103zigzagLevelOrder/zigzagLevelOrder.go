package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := treenode.TreeNode{Val: 1}
	root.Left = &treenode.TreeNode{Val: 2}
	root.Right = &treenode.TreeNode{Val: 3}
	root.Left.Left = &treenode.TreeNode{Val: 4}
	root.Right.Right = &treenode.TreeNode{Val: 5}
	res := zigzagLevelOrder2(&root)
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
func zigzagLevelOrder(root *treenode.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	curLevel := make([]*treenode.TreeNode, 0)
	curLevel = append(curLevel, root)
	left := true // 表示从左开始
	for len(curLevel) > 0 {
		s := make([]int, len(curLevel))
		nextLevel := make([]*treenode.TreeNode, 0)
		for i, n := range curLevel {
			if left {
				s[i] = n.Val
			} else {
				s[len(curLevel)-i-1] = n.Val
			}
			if n.Left != nil {
				nextLevel = append(nextLevel, n.Left)
			}
			if n.Right != nil {
				nextLevel = append(nextLevel, n.Right)
			}
		}
		left = !left
		res = append(res, s)
		curLevel = nextLevel
	}
	return res
}

func zigzagLevelOrder2(root *treenode.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	zig := false
	for len(queue) > 0 {
		thislevle := make([]*treenode.TreeNode, 0)
		ans := make([]int, 0)
		lenght := len(queue)
		for i := 0; i < lenght; i++ {
			var first *treenode.TreeNode
			if zig {
				first = queue[0]
				queue = queue[1:]
			} else {
				first = queue[len(queue)-1]
				queue = queue[:len(queue)-1]
			}
			ans = append(ans, first.Val)
			if first.Left != nil {
				thislevle = append(thislevle, first.Left)
			}
			if first.Right != nil {
				thislevle = append(thislevle, first.Right)
			}
		}
		zig = !zig
		queue = thislevle
		res = append(res, ans)
	}
	return res
}
