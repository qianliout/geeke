package main

import (
	"sort"

	"outback/leetcode/back/common/treenode"
)

func recoverTree2(root *treenode.TreeNode) {
	stark := make([]*treenode.TreeNode, 0)
	var (
		x   *treenode.TreeNode
		y   *treenode.TreeNode
		pre *treenode.TreeNode
	)

	for len(stark) > 0 || root != nil {
		for root != nil {
			stark = append(stark, root)
			root = root.Left
		}
		root := stark[len(stark)-1]
		stark = stark[:len(stark)-1]

		if pre != nil && pre.Val < root.Val {
			y = root
			if x == nil {
				x = pre
			} else {
				break
			}
		}

		pre = root
		root = root.Right
	}

	x.Val, y.Val = y.Val, x.Val
}

var (
	pre *treenode.TreeNode
	x   *treenode.TreeNode
	y   *treenode.TreeNode
)

func findTwo(root *treenode.TreeNode) {
	if root == nil {
		return
	}
	findTwo(root.Left)
	if pre != nil && root.Val < pre.Val {
		if x != nil {
			y = root
			return
		}
		x, y = pre, root
	}
	pre = root
	findTwo(root.Right)
}

// 直接中序遍历，再填回去
//方法一：直接中序
func inorderAppend(root *treenode.TreeNode, array *[]int) {
	if root == nil {
		return
	}
	inorderAppend(root.Left, array)
	*array = append(*array, root.Val)
	inorderAppend(root.Right, array)
}

func inorderFill(root *treenode.TreeNode, array []int, current *int) {
	if root == nil {
		return
	}
	inorderFill(root.Left, array, current)
	root.Val = array[*current]
	*current++
	inorderFill(root.Right, array, current)
}

func recoverTree3(root *treenode.TreeNode) {
	array := make([]int, 0)
	inorderAppend(root, &array)
	sort.Ints(array)
	iterator := 0
	inorderFill(root, array, &iterator)
}
