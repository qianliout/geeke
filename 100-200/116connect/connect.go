package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 7)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
*/
func connect(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
		root.Right.Next = getNextNode(root)
	}
	if root.Left == nil && root.Right != nil {
		root.Right.Next = getNextNode(root)
	}
	if root.Right == nil && root.Left != nil {
		root.Left.Next = getNextNode(root)
	}
	// 这里要注意：先递归右子树，否则右子树根节点next关系没建立好，左子树到右子树子节点无法正确挂载
	root.Right = connect(root.Right)
	root.Left = connect(root.Left)
	return root
}

func getNextNode(root *Node) *Node {
	for root.Next != nil {
		if root.Next.Left != nil {
			return root.Next.Left
		}
		if root.Next.Right != nil {
			return root.Next.Right
		}
		root = root.Next
	}
	return nil
}
