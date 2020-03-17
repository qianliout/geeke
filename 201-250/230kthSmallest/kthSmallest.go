package main

import (
	"fmt"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	res := recursion(root, 1)
	fmt.Println("res is ", res)
}

/*
给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。
示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？
链接：https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
*/
func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	inOrder(root, &res, k)
	if len(res) >= k {
		return res[k-1]
	}
	return 0
}

// 中序遍历()
func inOrder(root *TreeNode, res *[]int, k int) {
	if len(*res) >= k {
		return
	}
	if root == nil || len(*res) >= k {
		return
	}

	if root != nil && root.Left != nil {
		inOrder(root.Left, res, k)
	}
	if root != nil {
		*res = append(*res, root.Val)
	}
	if root != nil && root.Right != nil {
		inOrder(root.Right, res, k)
	}
}

// 因为是二叉搜索树，所以可以使用递归的方式
func recursion(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	leftCount := countNodes(root.Left)
	if leftCount+1 == k {
		return root.Val
	} else if leftCount+1 < k {
		return recursion(root.Right, k-leftCount-1)
	} else {
		return recursion(root.Left, k)
	}
}

func countNodes(root *TreeNode) int {
	num := 0
	if root != nil {
		num++
	}
	if root != nil && root.Left != nil {
		num += countNodes(root.Left)
	}
	if root != nil && root.Right != nil {
		num += countNodes(root.Right)
	}
	return num
}
