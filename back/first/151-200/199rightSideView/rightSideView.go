package main

import (
	"fmt"

	"qianliout/leetcode/back/common/treenode"
)

func main() {
	root := &treenode.TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	// root.Left.Right = &TreeNode{Val: 5}
	// root.Right = &TreeNode{Val: 3}
	// root.Right.Right = &TreeNode{Val: 4}
	res := rightSideView2(root)
	fmt.Println("res is ", res)
}

/*
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
示例:
输入: [1,2,3,null,5,null,4]
输出: [1, 3, 4]
解释:
   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
*/

// bfs
func rightSideView(root *treenode.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	res = append(res, root.Val)
	for len(queue) > 0 {
		thisLevel := make([]*treenode.TreeNode, 0)
		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:len(queue)]
			if first.Left != nil {
				thisLevel = append(thisLevel, first.Left)
			}
			if first.Right != nil {
				thisLevel = append(thisLevel, first.Right)
			}
		}
		if len(thisLevel) > 0 {
			res = append(res, thisLevel[len(thisLevel)-1].Val)
			queue = thisLevel
		} else {
			break
		}
	}
	return res
}

// dfs
var depth int

func rightSideView2(root *treenode.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	depth = 0 // 在我自己的环境中，是对的，但是在提交时，这里最好初始化
	dfs(root, &res, 0)
	return res
}

// 这道题的dfs一定要好好理解，关键点就是这个depth控制
func dfs(root *treenode.TreeNode, res *[]int, now int) {
	if root == nil {
		return
	}
	if now == depth {
		*res = append(*res, root.Val)
		depth++
	}
	dfs(root.Right, res, now+1)
	dfs(root.Left, res, now+1)
}
