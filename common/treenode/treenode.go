package treenode

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToTreeNode(nums []interface{}) *TreeNode {
	treeNode := &TreeNode{}
	return treeNode
}

func PreOrderTraversal(root *TreeNode) {
	if root != nil {
		traversePath = append(traversePath, root.Value)
		PreOrderTraversal(root.Left)
		fmt.Println(root.Value)
		PreOrderTraversal(root.Right)
	}
}

func InOrderTraversal(root *TreeNode) {
	if root != nil {
		PreOrderTraversal(root.Left)
		traversePath = append(traversePath, root.Value)
		fmt.Println(root.Value)
		PreOrderTraversal(root.Right)
	}
}

func PostOrderTraversal(root *TreeNode) {
	if root != nil {
		PreOrderTraversal(root.Right)
		//traversePath = append(traversePath, root.Value)
		fmt.Println(root.Value)
		PreOrderTraversal(root.Left)
	}
}
