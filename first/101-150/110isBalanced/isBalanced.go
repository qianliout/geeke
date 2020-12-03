package main

import (
	"fmt"
	"math"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	// root.Right = &TreeNode{Val: 5}
	// root.Right.Right = &TreeNode{Val: 8}

	res := dept(&root, 0)
	fmt.Println("res is ", res)
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := dept(root.Left, 0)
	right := dept(root.Right, 0)
	if math.Abs(float64(left-right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func dept(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	left := dept(root.Left, depth+1)
	right := dept(root.Right, depth+1)
	if left > right {
		return left
	}
	return right
}
