package bfstfs

import "outback/leetcode/common/treenode"

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
*/

func levelOrder(root *treenode.TreeNode) [][]int {

}

func levelOrderBfs(root *treenode.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := make([]*treenode.TreeNode, 0)

	for len(queue) > 0 {
		level_size = len(queue)
		current := make([]int, 0)
		for _, node := range queue {

			current=append(current,node.Val)

		}


	}

}

func levelOrderDfs(node treenode.TreeNode) [][]int {

}
