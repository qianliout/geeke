package main

import (
	"fmt"
	"math"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := treenode.TreeNode{Val: 2}
	root.Left = &treenode.TreeNode{Val: 1}
	root.Right = &treenode.TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 3}
	// root.Right.Right = &TreeNode{Val: 6}
	res := isValidBST4(&root)
	fmt.Println("res is ", res)
}

/*
   给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
    节点的左子树只包含小于当前节点的数。
    节点的右子树只包含大于当前节点的数。
    所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
*/

func isValidBST(root *treenode.TreeNode) bool {
	res := make([]int, 0)
	inorder(root, &res)
	if len(res) <= 1 {
		return true
	}
	for i := 1; i <= len(res)-1; i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

// 中序遍历一定是一个升序数组，可以使用这个特性,
// 我们在中序遍历的时候实时检查当前节点的值是否大于前一个中序遍历到的节点的值即可
func inorder(root *treenode.TreeNode, res *[]int) {
	if root != nil {
		inorder(root.Left, res)
		*res = append(*res, root.Val)
		inorder(root.Right, res)
	}
}

func isValidBST4(root *treenode.TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *treenode.TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

// 递归的方法,记录的是最大最小的结点
func IsValideByRecusion(root, min, max *treenode.TreeNode) bool {
	// min 记录的是当前最小的那个结点，max记录的当前最大的那个结点，这是关键，不是左结点和右结点
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return IsValideByRecusion(root.Left, min, root) && IsValideByRecusion(root.Right, root, max)
}
