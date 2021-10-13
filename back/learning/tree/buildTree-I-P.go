package tree

import (
	treenode2 "outback/leetcode/back/common/treenode"
)

/*
根据一棵树的中序遍历与后序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

func BuildTreeInorderAndPostOrder(inorder []int, postorder []int) *treenode2.TreeNode {
	inpos := make(map[int]int, 0)
	for key, value := range inorder {
		inpos[value] = key
	}
	length := len(inorder)
	return BuildTreeHelper(postorder, inpos, length-1, length-1, 0)
}

func BuildTreeHelper(postorder []int, inpos map[int]int, inEnd, postEnd, postStart int) *treenode2.TreeNode {
	if postEnd < postStart {
		return nil
	}
	root := treenode2.TreeNode{Val: postorder[postEnd]}
	rootIndex := inpos[postorder[postEnd]]
	rightLen := inEnd - rootIndex

	// build right
	root.Right = BuildTreeHelper(postorder, inpos, inEnd, postEnd-1, postEnd-rightLen-1)
	root.Left = BuildTreeHelper(postorder, inpos, rootIndex-1, postEnd-rightLen-1, postStart)
	return &root
}
