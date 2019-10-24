package dfsbfs

import (
	"fmt"
	"outback/leetcode/common/treenode"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := BuildTree1(preorder, inorder)
	n := treenode.PreOrderTraversal(root)
	fmt.Println(n)
}
