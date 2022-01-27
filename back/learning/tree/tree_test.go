package tree

import (
	"fmt"
	"testing"

	treenode2 "qianliout/leetcode/back/common/treenode"
)

func TestBuildTreeInorderAndPostOrder(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := BuildTreeInorderAndPostOrder(inorder, postorder)
	treenode2.PostOrderTraversal(root)
}

func TestBuildTreeInorderAndPostOrder105(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := BuildTree105(inorder, postorder)
	treenode2.PostOrderTraversal(root)
}

func TestIsValidBST(t *testing.T) {
	root := treenode2.TreeNode{Val: 2}
	root.Left = &treenode2.TreeNode{Val: 1}
	root.Right = &treenode2.TreeNode{Val: 3}
	root.Right.Left = &treenode2.TreeNode{Val: 2}
	n := IsValidBST(&root)
	fmt.Println(n)
}
