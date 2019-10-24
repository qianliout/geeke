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
	res := make([]int, 0)
	PostOrderTraversal(&root, res)
	fmt.Println(res)

}
