package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {

}

/*
给定一个二叉树，根节点为第1层，深度为 1。在其第 d 层追加一行值为 v 的节点。

添加规则：给定一个深度值 d （正整数），针对深度为 d-1 层的每一非空节点 N，为 N 创建两个值为 v 的左子树和右子树。
将 N 原先的左子树，连接为新节点 v 的左子树；将 N 原先的右子树，连接为新节点 v 的右子树。
如果 d 的值为 1，深度 d - 1 不存在，则创建一个新的根节点 v，原先的整棵树将作为 v 的左子树。
示例 1:
输入:
二叉树如下所示:
       4
     /   \
    2     6
   / \   /
  3   1 5
v = 1
d = 2
输出:
       4
      / \
     1   1
    /     \
   2       6
  / \     /
 3   1   5
示例 2:
输入:
二叉树如下所示:
      4
     /
    2
   / \
  3   1
v = 1
d = 3
输出:
      4
     /
    2
   / \
  1   1
 /     \
3       1
注意:
输入的深度值 d 的范围是：[1，二叉树最大深度 + 1]。
输入的二叉树至少有一个节点。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *treenode.TreeNode, v int, d int) *treenode.TreeNode {
	queue := make([]*treenode.TreeNode, 0)
	queue = append(queue, root)
	flag := false
	for len(queue) > 0 {
		if d == 2 { // 这个数字要好好确定
			flag = true
			for i := 0; i < len(queue); i++ {
				node := queue[i]
				left := node.Left
				right := node.Right

				node.Left = &treenode.TreeNode{Val: v}
				node.Right = &treenode.TreeNode{Val: v}
				node.Left.Left = left
				node.Right.Right = right
			}
			// 操作了
			break
		}
		thisLeve := make([]*treenode.TreeNode, 0)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				thisLeve = append(thisLeve, queue[i].Left)
			}
			if queue[i].Right != nil {
				thisLeve = append(thisLeve, queue[i].Right)
			}
		}
		d--
		queue = thisLeve
	}
	if !flag {
		node := &treenode.TreeNode{Val: v}
		node.Left = root
		return node
	}

	return root
}
