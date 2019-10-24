package treenode

import (
	"fmt"
	"testing"
)

func TestPostOrderTraversal(t *testing.T) {
	root := TreeNode{Value: 3}
	root.Left = &TreeNode{Value: 4}
	root.Right = &TreeNode{Value: 5}
	root.Left.Right = &TreeNode{Value: 6}
	root.Left.Left = &TreeNode{Value: 7}
	root.Right.Left = &TreeNode{Value: 8}
	root.Right.Right = &TreeNode{Value: 9}
	res := make([]int, 0)
	InOrderTraversal(&root)
	fmt.Println(res)

}
