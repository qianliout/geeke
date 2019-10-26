package treenode

import (
	"testing"
)

func newTree() *TreeNode {
	/*
		返回如下的二叉树：
		    3
		   / \
		  9  20
		    /  \
		   15   7
	*/

	root := TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return &root
}

func TestPostOrderTraversal(t *testing.T) {
	root := newTree()
	InOrderTraversal(root)
	PreOrderTraversal(root)
	PostOrderTraversal(root)
}


