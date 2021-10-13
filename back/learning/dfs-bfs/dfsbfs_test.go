package dfsbfs

import (
	"testing"

	treenode2 "outback/leetcode/back/common/treenode"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTreePreAndIn(preorder, inorder)
	treenode2.PostOrderTraversal(root)
}

func TestInAndPostHelper(t *testing.T) {
	preorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTreeInAndPost(preorder, inorder)
	treenode2.InOrderTraversal(root)
	//treenode.PostOrderTraversal(root)
}
