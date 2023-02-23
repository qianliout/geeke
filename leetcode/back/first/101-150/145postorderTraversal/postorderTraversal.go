package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := &treenode.TreeNode{Val: 1}
	root.Left = &treenode.TreeNode{Val: 4}
	root.Left.Left = &treenode.TreeNode{Val: 2}
	root.Left.Right = &treenode.TreeNode{Val: 5}
	root.Right = &treenode.TreeNode{Val: 3}
	res := postorderTraversal(root)
	fmt.Println("res is ", res)
}

/*
给定一个二叉树，返回它的 后序 遍历。
示例:
输入: [1,null,2,3]
   1
    \
     2
    /
   3
输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
*/
func postorderTraversal(root *treenode.TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	//post(root, &res)
	Iter(root, &res)
	return res
}

func post(root *treenode.TreeNode, res *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		post(root.Left, res)
	}
	if root.Right != nil {
		post(root.Right, res)
	}
	*res = append(*res, root.Val)
}

func Iter(root *treenode.TreeNode, res *[]int) {
	stack1 := make([]*treenode.TreeNode, 0)
	stack2 := make([]*treenode.TreeNode, 0)
	stack1 = append(stack1, root)

	for len(stack1) > 0 {
		node := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, node)
		//这里压栈的顺序很重要
		if node.Left != nil {
			stack1 = append(stack1, node.Left)
		}
		if node.Right != nil {
			stack1 = append(stack1, node.Right)
		}
	}
	for len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]
		*res = append(*res, node.Val)
	}
}
