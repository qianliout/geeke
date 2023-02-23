package main

import (
	"outback/leetcode/back/common/treenode"
)

func main() {

}

/*
给出一个完全二叉树，求出该树的节点个数。
说明：
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
示例:
输入:
    1
   / \
  2   3
 / \  /
4  5 6
输出: 6
链接：https://leetcode-cn.com/problems/count-complete-tree-nodes
*/
// 最简单的方法就是使用递归遍历一次这棵树,但是没有用到满三叉树的特性
func countNodes(root *treenode.TreeNode) int {
	num := 0
	if root != nil {
		num++
	}
	if root != nil && root.Left != nil {
		num += countNodes(root.Left)
	}
	if root != nil && root.Right != nil {
		num += countNodes(root.Right)
	}
	return num
}

func countNodes2(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDept := countDept(root.Left)
	rightDept := countDept(root.Right)
	if leftDept == rightDept {
		// 左右层级一样,说左树一定是满二叉树,左树的节点个数是2^n-1 加上root节点刚好 2^n
		return countNodes2(root.Right) + (1 << leftDept)
	} else {
		// 说明左树可能不满,但是到数第层,也就是右树一定是一个满树
		return countNodes2(root.Left) + (1 << rightDept)
	}
}

func countDept(root *treenode.TreeNode) int {
	num := 0
	if root != nil {
		num++
	}
	if root != nil && root.Left != nil {
		num += countDept(root.Left)
	}
	return num
}

func countDept2(root *treenode.TreeNode) int {
	num := 0
	for root != nil {
		num++
		root = root.Left
	}
	return num
}
