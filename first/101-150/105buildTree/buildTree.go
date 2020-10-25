package main

import (
	. "outback/leetcode/common/treenode"
)

func main() {

}

// 根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
// 你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
//
// 返回如下的二叉树：
//
//    3
//   / \
//  9  20
//    /  \
//   15   7

// 这是这道题最简单的理解方法
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 前序遍历第一个值为根节点
	root := TreeNode{Val: preorder[0]}
	// 因为没有重复元素，所以可以直接根据值来查找根节点在中序遍历中的位置
	mid := FindIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return &root
}

func FindIndex(list []int, value int) int {
	for key, v := range list {
		if v == value {
			return key
		}
	}
	return -1
}
