package main

import (
	"qianliout/leetcode/back/common/treenode"
)

func main() {

}

/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
*/

func zigzagLevelOrder(root *treenode.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	zig := true

	for len(queue) > 0 {
		lenght := len(queue)
		nex := make([]int, lenght)
		nexnode := make([]*treenode.TreeNode, 0)
		for i := 0; i < lenght; i++ {
			first := queue[0]
			queue = queue[1:]
			if zig {
				nex[i] = first.Val
			} else {
				nex[lenght-i-1] = first.Val
			}

			if first.Left != nil {
				nexnode = append(nexnode, first.Left)
			}
			if first.Right != nil {
				nexnode = append(nexnode, first.Right)
			}
		}
		zig = !zig
		res = append(res, nex)
		queue = nexnode
	}
	return res
}
