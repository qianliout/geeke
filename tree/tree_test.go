package tree

import (
	"outback/leetcode/common/treenode"
	"testing"
)

func TestBuildTreeInorderAndPostOrder(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := BuildTreeInorderAndPostOrder(inorder, postorder)
	treenode.PostOrderTraversal(root)
}

func TestBuildTreeInorderAndPostOrder105(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := BuildTree105(inorder, postorder)
	treenode.PostOrderTraversal(root)
}
