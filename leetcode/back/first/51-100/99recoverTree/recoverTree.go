package main

import (
	"fmt"

	"outback/leetcode/back/common/treenode"
)

func main() {
	root := &treenode.TreeNode{Val: 3}
	root.Left = &treenode.TreeNode{Val: 1}
	//root.Left.Right = &TreeNode{Val: 2}
	root.Right = &treenode.TreeNode{Val: 4}
	root.Right.Left = &treenode.TreeNode{Val: 2}
	recoverTree2(root)
	treenode.InOrderTraversal(root)
	fmt.Println()
	treenode.PreOrderTraversal(root)
}

/*
二叉搜索树中的两个节点被错误地交换。
请在不改变其结构的情况下，恢复这棵树。
示例 1:
输入: [1,3,null,null,2]
   1
  /
 3
  \
   2
输出: [3,1,null,null,2]
   3
  /
 1
  \
   2
示例 2:
输入: [3,1,4,null,null,2]
  3
 / \
1   4
   /
  2
输出: [2,1,4,null,null,3]
  2
 / \
1   4
   /
  3
进阶:

    使用 O(n) 空间复杂度的解法很容易实现。
    你能想出一个只使用常数空间的解决方案吗？
链接：https://leetcode-cn.com/problems/recover-binary-search-tree
*/

// 我在本地一直运行都能获得正确答案，但是提交显示不成功
func recoverTree(root *treenode.TreeNode) {
	if root == nil {
		return
	}

	//inorderToSlice(root)
	//fmt.Println(res)
	//f, s := findSWap(res)
	//fmt.Println(f, s)
	//inorder(root, 2, f, s)
	//inorder(root, res)
	findTwo(root)
	x.Val, y.Val = y.Val, x.Val
}

func findSWap(s []int) (int, int) {
	n := len(s)
	fb := false
	first, second := 0, 0
	for i := 0; i < n; i++ {
		if i+1 < n && s[i] > s[i+1] {
			if !fb {
				first = s[i]
				second = s[i+1]
				fb = true
			} else {
				second = s[i+1]
				break
			}
		}
	}
	return first, second
}

var res []int

func inorderToSlice(root *treenode.TreeNode) {
	if res == nil {
		res = make([]int, 0)
	}
	if root == nil {
		return
	}
	inorderToSlice(root.Left)
	res = append(res, root.Val)
	inorderToSlice(root.Right)
}

func inorder(root *treenode.TreeNode, cont, fist, second int) {
	if root == nil {
		return
	}
	if root.Val == fist || root.Val == second {
		if root.Val == fist {
			root.Val = second
		} else {
			root.Val = fist
		}
		cont--
		if cont < 0 {
			return
		}
	}
	inorder(root.Left, cont, fist, second)
	inorder(root.Right, cont, fist, second)
}
