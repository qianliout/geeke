package main

import (
	"fmt"
	"math"

	. "outback/leetcode/common/treenode"
)

func main() {
	root := TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	//root.Right.Left = &TreeNode{Val: 3}
	//root.Right.Right = &TreeNode{Val: 6}
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

func isValidBST(root *TreeNode) bool {
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

//中序遍历一定是一个升序数组，可以使用这个特性
func inorder(root *TreeNode, res *[]int) {
	if root != nil {
		inorder(root.Left, res)
		*res = append(*res, root.Val)
		inorder(root.Right, res)
	}
}

// 上面的方法，需要遍历后再次进行遍历，todo 是否可以一次做完
func isValidBST2(root *TreeNode) bool {
	return inorder2(root, nil)
}

// todo 这里把pre定义成全局变量就对，why
func inorder2(root, pre *TreeNode) bool {
	fmt.Println(pre, root)
	if root == nil {
		return true
	}
	if !inorder2(root.Left, pre) { // 这里一定要注意的是，如果左边为真，还不能判定为真，但是左边为假，一定是假
		return false
	}

	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	// 为什么这里可以直接返回呢，因为程序走到这里，说明左边为真了。
	return inorder2(root.Right, pre)
}
func isValidBST3(root *TreeNode) bool {
	if root != nil {
		return true
	}
	return false
}

func isValidBST4(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min, max int) bool {
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
