package main

import (
	"outback/leetcode/back/common"
	"outback/leetcode/back/common/treenode"
)

/*
给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。
这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。
每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
示例 1:
输入:
           1
         /   \
        3     2
       / \     \
      5   3     9
输出: 4
解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。
*/
func widthOfBinaryTree(root *treenode.TreeNode) int {
	return bfs(root)
}

type item struct {
	node *treenode.TreeNode
	pas  int
}

func bfs(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	width := 0
	queue := make([]*item, 0)
	queue = append(queue, &item{
		node: root,
		pas:  0,
	})

	for len(queue) > 0 {
		thisLevle := make([]*item, 0)
		for _, it := range queue {
			if it.node.Left != nil {
				thisLevle = append(thisLevle, &item{
					node: it.node.Left,
					pas:  it.pas * 2,
				})
			}
			if it.node.Right != nil {
				thisLevle = append(thisLevle, &item{
					node: it.node.Right,
					pas:  it.pas*2 + 1,
				})
			}
		}
		width = common.Max(width, queue[len(queue)-1].pas-queue[0].pas+1)
		queue = thisLevle
	}
	return width
}
