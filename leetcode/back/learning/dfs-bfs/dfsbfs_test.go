package dfsbfs

import (
	"testing"

	"qianliout/leetcode/leetcode/common/treenode"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTreePreAndIn(preorder, inorder)
	treenode.PostOrderTraversal(root)
}

func TestInAndPostHelper(t *testing.T) {
	preorder := []int{9, 15, 7, 20, 3}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTreeInAndPost(preorder, inorder)
	treenode.InOrderTraversal(root)
	// treenode.PostOrderTraversal(root)
}
