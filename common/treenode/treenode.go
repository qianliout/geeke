package treenode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SliceToTreeNode(nums []interface{}) *TreeNode {
	treeNode := &TreeNode{}
	return treeNode
}

func PreOrderTraversal(root *TreeNode) {
	if root != nil {
		fmt.Printf("%d ", root.Val)
		PreOrderTraversal(root.Left)
		PreOrderTraversal(root.Right)
	}
}

// 中序排列
func InOrderTraversal(root *TreeNode) {
	if root != nil {
		InOrderTraversal(root.Left)
		fmt.Printf("%d ", root.Val)
		InOrderTraversal(root.Right)
	}
}

// 后序排列
func PostOrderTraversal(root *TreeNode) {
	if root != nil {
		InOrderTraversal(root.Right)
		InOrderTraversal(root.Left)
		fmt.Printf("%d ", root.Val)
	}
}

func InorderTraversalToSlice(root *TreeNode) []int {
	reslut := make([]int, 0)
	func(root *TreeNode) {
		if root != nil {
			InorderTraversalToSlice(root.Left)
			reslut = append(reslut, root.Val)
			InorderTraversalToSlice(root.Right)
		}
	}(root)
	return reslut
}
