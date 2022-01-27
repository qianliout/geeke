package dfsbfs

import (
	treenode2 "qianliout/leetcode/back/common/treenode"
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

func BuildTreeInAndPost(inorder, postorder []int) *treenode2.TreeNode {
	pos := make(map[int]int)
	for key, value := range inorder {
		pos[value] = key
	}
	return InAndPostHelper(postorder, len(postorder)-1, 0, 0, pos)

}

func InAndPostHelper(postorder []int, postStart, postEnd, inStart int, pos map[int]int) *treenode2.TreeNode {
	// 从后往前走
	if postEnd > postStart || postStart < 0 {
		return nil
	}
	root := treenode2.TreeNode{Val: postorder[postStart]}
	rootIndex := pos[postorder[postStart]]
	rightLen := postStart - rootIndex
	root.Right = InAndPostHelper(postorder, postStart-1, postStart-rightLen, rootIndex-1, pos)
	root.Left = InAndPostHelper(postorder, postStart-rightLen-1, postEnd, inStart, pos)
	return &root
}
